package gapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	pb "github.com/tunvx/simplebank/grpc/pb/moneytransfer"
	"github.com/tunvx/simplebank/moneytransfersrv/val"
	worker "github.com/tunvx/simplebank/moneytransfersrv/worker/kafka"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Customer use this api
func (service *Service) InternalTransferMoney(ctx context.Context, req *pb.InternalTransferMoneyRequest) (*pb.InternalTransferMoneyResponse, error) {
	// 1. Validate params and authorize user
	violations := validateFastInternalTransfer(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}
	if req.GetSrcAccNumber() == req.GetBeneAccNumber() {
		return nil, status.Errorf(codes.PermissionDenied, "transfer failed: source and beneficiary accounts cannot be the same.")
	}
	authPayload, err := service.authorizeUser(ctx, []string{util.CustomerRole, util.BankerRole})
	if err != nil {
		return nil, err
	}

	// 2. Get two account info for validate

	// validate source: status and currency
	sourceAccount, err := service.checkAccount(ctx, req.GetSrcAccNumber(), req.GetCurrencyType())
	if err != nil {
		return nil, err
	}

	if authPayload.Role == util.CustomerRole && authPayload.UserID != sourceAccount.CustomerID {
		err = status.Errorf(codes.PermissionDenied, "user_id cannot match with account_owner_id")
		return nil, err
	}

	// validate beneficiary account: status and currency
	beneficiaryAccount, err := service.checkAccount(ctx, req.GetBeneAccNumber(), req.GetCurrencyType())
	if err != nil {
		return nil, err
	}

	// 3. Prepare params
	sendingTranID, err := uuid.NewV7() // Create sending transaction id
	if err != nil {
		err = status.Errorf(codes.Internal, "error when create sending transaction Id: %s", err)
		return nil, err
	}
	receivingTranID, err := uuid.NewV7() // Create reveiving transaction id
	if err != nil {
		err = status.Errorf(codes.Internal, "error when create reveiving transaction Id: %s", err)
		return nil, err
	}
	sendingDescription := fmt.Sprintf("Noi dung: %s. CT tu %s %s toi %s %s tai MYBANK", req.GetMessage(),
		req.GetSrcAccNumber(), sourceAccount.OwnerName, req.GetBeneAccNumber(), beneficiaryAccount.OwnerName)
	receivingDescription := fmt.Sprintf("%s chuyen tien", beneficiaryAccount.OwnerName)

	// 4. Do transaction
	// Principle: avoid locking database resources for long periods of time
	// Question: But how to ensure eventual consistency?

	// 4.1 Transaction in-shard.
	if sourceAccount.ShardId == beneficiaryAccount.ShardId {
		arg := db.CreateCompleteTxInShardParams{
			SendingTransactionID:   sendingTranID,
			ReceivingTransactionID: receivingTranID,
			Amount:                 req.GetAmount(),
			SourceAccID:            sourceAccount.AccountID,
			BeneficiaryAccID:       beneficiaryAccount.AccountID,
			SendingDescription:     sendingDescription,
			ReceivingDescription:   receivingDescription,
		}

		shardID := sourceAccount.ShardId - 1

		result, err := service.stores[shardID].CreateCompleteTxInShard(ctx, arg)
		if err != nil {
			err = status.Errorf(codes.Internal, "error when performing transaction in-shard ( %d ): %s", shardID, err)
			return nil, err
		}

		response := &pb.InternalTransferMoneyResponse{
			IsSuccessful:         true,
			SourceAccount:        convertAccount(result.SourceAccount),
			SendingTransaction:   convertAccountTransaction(result.SendingTransaction),
			BeneficiaryAccount:   convertAccount(result.BeneficiaryAccount),
			ReceivingTransaction: convertAccountTransaction(result.ReceivingTransaction),
		}
		return response, nil
	}

	// 4.2 Transaction between shards
	shardID1 := sourceAccount.ShardId - 1
	shardID2 := beneficiaryAccount.ShardId - 1
	log.Debug().Msgf("MoneyTransfer Service: Account ( %d ) has shard_id ( %d )", sourceAccount.AccountID, shardID1)
	log.Debug().Msgf("MoneyTransfer Service: Account ( %d ) has shard_id ( %d )", beneficiaryAccount.AccountID, shardID2)

	// 4.2.1 Send event to kafka
	payload := worker.PayloadInternalTransferMoney{
		SendingTranID:   sendingTranID,
		ReceivingTranID: receivingTranID,
		Amount:          req.GetAmount(),
		CurrencyType:    req.GetCurrencyType(),
		SrcAccNumber:    req.GetSrcAccNumber(),
		SrcAccShardId:   sourceAccount.ShardId,
		BeneAccNumber:   req.GetBeneAccNumber(),
		BeneAccShardId:  beneficiaryAccount.ShardId,
	}
	err = service.taskProducer.SendTaskInternalTransferMoney(ctx, &payload)
	if err != nil {
		log.Error().Msgf("Error when sending task to Kafka: %v", err)
	}

	// 4.2.2 Create Sending transaction
	sendingTxArg := db.CreateTxTransferMoneyParams{
		TransactionID:     sendingTranID,
		Amount:            req.GetAmount(),
		SourceAccID:       sourceAccount.AccountID,
		Description:       sendingDescription,
		TransactionType:   db.TransactiontypeInternalSend,
		TransactionStatus: db.TransactionstatusCompleted,
	}
	sendingTxResult, err := service.stores[shardID1].CreateTxTransferMoney(ctx, sendingTxArg)
	if err != nil {
		err = status.Errorf(codes.Internal, "error in phase transfer money: %s", err)
		return nil, err
	}

	// 4.2.3 Create Receiving transaction
	receivingTxArg := db.CreateTxReceiveMoneyParams{
		TransactionID:     receivingTranID,
		Amount:            req.GetAmount(),
		BeneficiaryAccID:  beneficiaryAccount.AccountID,
		Description:       receivingDescription,
		TransactionType:   db.TransactiontypeInternalReceive,
		TransactionStatus: db.TransactionstatusCompleted,
	}
	receivingTxResult, err := service.stores[shardID2].CreateTxReceiveMoney(ctx, receivingTxArg)
	if err != nil {
		err = status.Errorf(codes.Internal, "error in phase receive money: %s", err)
		return nil, err
	}

	response := &pb.InternalTransferMoneyResponse{
		IsSuccessful:         true,
		SourceAccount:        convertAccount(sendingTxResult.SourceAccount),
		SendingTransaction:   convertAccountTransaction(sendingTxResult.SendingTransaction),
		BeneficiaryAccount:   convertAccount(receivingTxResult.BeneficiaryAccount),
		ReceivingTransaction: convertAccountTransaction(receivingTxResult.ReceivingTransaction),
	}
	return response, nil
}

func validateFastInternalTransfer(req *pb.InternalTransferMoneyRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Customer Real ID
	if err := val.ValidateTransferAmount(req.GetAmount(), req.GetCurrencyType()); err != nil {
		violations = append(violations, errga.FieldViolation("amount", err))
	}

	if err := val.ValidateCurrency(req.GetCurrencyType()); err != nil {
		violations = append(violations, errga.FieldViolation("currency_type", err))
	}

	if err := val.ValidateAccountNumber(req.GetSrcAccNumber()); err != nil {
		violations = append(violations, errga.FieldViolation("sender_acc_number", err))
	}

	if err := val.ValidateAccountNumber(req.GetBeneAccNumber()); err != nil {
		violations = append(violations, errga.FieldViolation("recipient_acc_number", err))
	}

	return violations
}
