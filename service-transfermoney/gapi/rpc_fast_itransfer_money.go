package gapi

import (
	"context"

	"github.com/rs/zerolog/log"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/transfermoney"
	db "github.com/tunvx/simplebank/management/db/sqlc"
	"github.com/tunvx/simplebank/transfermoney/gapi/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) FastInternalTransfer(ctx context.Context, req *pb.FastInternalTransferRequest) (*pb.FastInternalTransferResponse, error) {
	authPayload, err := service.authorizeUser(ctx, []string{util.BankerRole, util.CustomerRole})
	if err != nil {
		return nil, err
	}

	userID := authPayload.UserID
	shardID := util.ExtractShardID(userID)

	// Log the start of the transfer request
	log.Info().
		Int64("user_id", authPayload.UserID).
		Str("session_id", authPayload.ID.String()).
		Str("operation", "FastInternalTransfer").
		Str("role", authPayload.Role).
		Int64("amount", req.GetAmount()).
		Str("currency", req.GetCurrencyType()).
		Str("sender_account", req.GetSenderAccNumber()).
		Str("recipient_account", req.GetRecipientAccNumber()).
		Msg("internal transfer initiated")

	violations := validateFastInternalTransfer(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	// Validate source: status and currency
	senderAccount, err := service.getAndValidateAccount(ctx, shardID, req.SenderAccNumber, req.CurrencyType)
	if err != nil {
		return nil, err
	}

	if authPayload.Role == util.CustomerRole && authPayload.UserID != senderAccount.CustomerID {
		err = status.Errorf(codes.PermissionDenied, "user_id cannot match with account_owner_id")
		return nil, err
	}

	// Validate beneficiary account: status and currency
	recipientAccount, err := service.getAndValidateAccount(ctx, shardID, req.RecipientAccNumber, req.CurrencyType)
	if err != nil {
		return nil, err
	}

	arg := db.CreateCompleteTxParams{
		Amount:             req.GetAmount(),
		SenderAccountID:    senderAccount.AccountID,
		RecipientAccountID: recipientAccount.AccountID,
		Message:            req.GetMessage(),

		// TODO: Push notification
		AfterTransfer: func(amount int64, serder_account db.Account, recipient_account db.Account) error {
			return nil
		},
	}

	result, err := service.stores[shardID].CreateCompleteTx(ctx, arg)
	if err != nil {
		err = status.Errorf(codes.Internal, "internal transfer error: %s", err)
		return nil, err
	}

	// Log the successful transfer
	log.Info().
		Int64("user_id", authPayload.UserID).
		Str("session_id", authPayload.ID.String()).
		Str("operation", "FastInternalTransfer").
		Str("role", authPayload.Role).
		Int64("amount", req.GetAmount()).
		Str("currency", req.GetCurrencyType()).
		Str("sender_account", req.GetSenderAccNumber()).
		Int64("sender_balance_after", result.SenderAccount.CurrentBalance).
		Str("recipient_account", req.GetRecipientAccNumber()).
		Int64("recipient_balance_after", result.RecipientAccount.CurrentBalance).
		Str("transaction_status", "completed").
		Msg("internal transfer completed successfully")

	response := &pb.FastInternalTransferResponse{
		SenderAccount:    convertAccount(result.SenderAccount),
		RecipientAccount: convertAccount(result.RecipientAccount),
	}
	return response, nil
}

func validateFastInternalTransfer(req *pb.FastInternalTransferRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Customer Real ID
	if err := val.ValidateTransferAmount(req.GetAmount(), req.GetCurrencyType()); err != nil {
		violations = append(violations, errga.FieldViolation("amount", err))
	}

	if err := val.ValidateAccountNumber(req.GetSenderAccNumber()); err != nil {
		violations = append(violations, errga.FieldViolation("sender_acc_number", err))
	}

	if err := val.ValidateAccountNumber(req.GetRecipientAccNumber()); err != nil {
		violations = append(violations, errga.FieldViolation("recipient_acc_number", err))
	}

	if err := val.ValidateFullName(req.GetRecipientName()); err != nil {
		violations = append(violations, errga.FieldViolation("recipient_name", err))
	}

	if err := val.ValidateCurrency(req.GetCurrencyType()); err != nil {
		violations = append(violations, errga.FieldViolation("currency_type", err))
	}

	return violations
}
