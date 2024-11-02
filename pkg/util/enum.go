package util

// TransactionType constants
const (
	RepayLoan      = "repay_loan"
	DepositSavings = "deposit_savings"
	SendMoney      = "send_money"
	ReceiveMoney   = "receive_money"
	DepositCash    = "deposit_cash"
	WithdrawCash   = "withdraw_cash"
	Others         = "others"
)

// IsSupportedTransactionType returns true if the transaction type is supported
func IsSupportedTransactionType(transactionType string) bool {
	switch transactionType {
	case RepayLoan, DepositSavings, SendMoney, ReceiveMoney, DepositCash, WithdrawCash, Others:
		return true
	}
	return false
}

// CustomerSegment constants
const (
	Individual       = "individual"
	SmallEnterprise  = "small_enterprise"
	MediumEnterprise = "medium_enterprise"
	LargeEnterprise  = "large_enterprise"
	Institutional    = "institutional"
)

// IsSupportedCustomerSegment returns true if the customer segment is supported
func IsSupportedCustomerSegment(segment string) bool {
	switch segment {
	case Individual, SmallEnterprise, MediumEnterprise, LargeEnterprise, Institutional:
		return true
	}
	return false
}

// CustomerTier constants
const (
	Regular  = "regular"
	Bronze   = "bronze"
	Silver   = "silver"
	Gold     = "gold"
	Platinum = "platinum"
	Diamond  = "diamond"
)

// IsSupportedCustomerTier returns true if the customer tier is supported
func IsSupportedCustomerTier(tier string) bool {
	switch tier {
	case Regular, Bronze, Silver, Gold, Platinum, Diamond:
		return true
	}
	return false
}

// FinancialStatus constants
const (
	Excellent = "excellent"
	VeryGood  = "very_good"
	Good      = "good"
	Fair      = "fair"
	Poor      = "poor"
	VeryPoor  = "very_poor"
)

// IsSupportedFinancialStatus returns true if the financial status is supported
func IsSupportedFinancialStatus(status string) bool {
	switch status {
	case Excellent, VeryGood, Good, Fair, Poor, VeryPoor:
		return true
	}
	return false
}

// EmployeePosition constants
const (
	BranchManager    = "branch_manager"
	Banker           = "banker"
	FinancialAnalyst = "financial_analyst"
)

// IsSupportedEmployeePosition returns true if the employee position is supported
func IsSupportedEmployeePosition(position string) bool {
	switch position {
	case BranchManager, Banker, FinancialAnalyst:
		return true
	}
	return false
}

// EmployeeStatus constants
const (
	Active   = "active"
	Inactive = "inactive"
)

// IsSupportedEmployeeStatus returns true if the employee status is supported
func IsSupportedEmployeeStatus(status string) bool {
	switch status {
	case Active, Inactive:
		return true
	}
	return false
}

// BankStatus constants
const (
	BankActive   = "active"
	BankInactive = "inactive"
)

// IsSupportedBankStatus returns true if the bank status is supported
func IsSupportedBankStatus(status string) bool {
	switch status {
	case BankActive, BankInactive:
		return true
	}
	return false
}

// BranchStatus constants
const (
	BranchActive   = "active"
	BranchInactive = "inactive"
)

// IsSupportedBranchStatus returns true if the branch status is supported
func IsSupportedBranchStatus(status string) bool {
	switch status {
	case BranchActive, BranchInactive:
		return true
	}
	return false
}

// AccountStatus constants
const (
	AccountActive   = "active"
	AccountInactive = "inactive"
)

// IsSupportedAccountStatus returns true if the account status is supported
func IsSupportedAccountStatus(status string) bool {
	switch status {
	case AccountActive, AccountInactive:
		return true
	}
	return false
}

// CurrencyType constants
const (
	VND = "VND"
	USD = "USD"
)

// IsSupportedCurrencyType returns true if the currency type is supported
func IsSupportedCurrencyType(currency string) bool {
	switch currency {
	case VND, USD:
		return true
	}
	return false
}

// MaturityInstruction constants
const (
	Reinvest       = "reinvest"
	InterestOnly   = "interest_only"
	FullWithdrawal = "full_withdrawal"
)

// IsSupportedMaturityInstruction returns true if the maturity instruction is supported
func IsSupportedMaturityInstruction(instruction string) bool {
	switch instruction {
	case Reinvest, InterestOnly, FullWithdrawal:
		return true
	}
	return false
}

// TransactionStatus constants
const (
	Pending   = "pending"
	Completed = "completed"
	Failed    = "failed"
)

// IsSupportedTransactionStatus returns true if the transaction status is supported
func IsSupportedTransactionStatus(status string) bool {
	switch status {
	case Pending, Completed, Failed:
		return true
	}
	return false
}

// SavingStatus constants
const (
	SavingActive     = "active"
	SavingCompleted  = "completed"
	SavingTerminated = "terminated"
	SavingInactive   = "inactive"
)

// IsSupportedSavingStatus returns true if the saving status is supported
func IsSupportedSavingStatus(status string) bool {
	switch status {
	case SavingActive, SavingCompleted, SavingTerminated, SavingInactive:
		return true
	}
	return false
}

// LoanStatus constants
const (
	Accruing      = "accruing"
	PaidOff       = "paid_off"
	Extended      = "extended"
	NonPerforming = "non_performing"
)

// IsSupportedLoanStatus returns true if the loan status is supported
func IsSupportedLoanStatus(status string) bool {
	switch status {
	case Accruing, PaidOff, Extended, NonPerforming:
		return true
	}
	return false
}
