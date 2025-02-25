package util

// 1. TransactionStatus constants
const (
	Completed = "completed"
	Pending   = "pending"
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

// 2. TransactionType constants
const (
	InternalSend    = "internal_send"
	InternalReceive = "internal_receive"
	ExternalSend    = "external_send"
	ExternalReceive = "external_receive"
	RepayLoan       = "repay_loan"
	DepositSavings  = "deposit_savings"
	Others          = "others"
)

// IsSupportedTransactionType returns true if the transaction type is supported
func IsSupportedTransactionType(transactionType string) bool {
	switch transactionType {
	case RepayLoan, DepositSavings, InternalSend, InternalReceive, ExternalSend, ExternalReceive, Others:
		return true
	}
	return false
}

// 3. CurrencyType constants
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

// 4. AccountStatus constants
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

// 5. CustomerTier constants
const (
	Standard = "standard"
	Bronze   = "bronze"
	Silver   = "silver"
	Gold     = "gold"
	Platinum = "platinum"
	Diamond  = "diamond"
	Vip      = "vip"
)

// IsSupportedCustomerTier returns true if the customer tier is supported
func IsSupportedCustomerTier(tier string) bool {
	switch tier {
	case Standard, Bronze, Silver, Gold, Platinum, Diamond, Vip:
		return true
	}
	return false
}

// 6. CustomerSegment constants
const (
	Retail        = "retail"
	SmallBusiness = "small_business"
	Corporate     = "corporate"
	Institutional = "institutional"
	Government    = "government"
)

// IsSupportedCustomerSegment returns true if the customer segment is supported
func IsSupportedCustomerSegment(segment string) bool {
	switch segment {
	case Retail, SmallBusiness, Corporate, Institutional, Government:
		return true
	}
	return false
}

// 7. FinancialStatus constants
const (
	VeryGood  = "very_good"
	Good      = "good"
	Average   = "average"
	LowRisk   = "low_risk"
	HighRisk  = "high_risk"
	Defaulted = "defaulted"
)

// IsSupportedFinancialStatus returns true if the financial status is supported
func IsSupportedFinancialStatus(status string) bool {
	switch status {
	case VeryGood, Good, Average, LowRisk, HighRisk, Defaulted:
		return true
	}
	return false
}

//*****************************************************************////*****************************************************************//
//*****************************************************************////*****************************************************************//
//*****************************************************************////*****************************************************************//
