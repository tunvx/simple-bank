package val

import (
	"fmt"
	"net/mail"
	"regexp"
	"time"
	"unicode"

	"github.com/tunvx/simplebank/pkg/util"
)

// Regex patterns for username and other validations
var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isPhoneNumber   = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`).MatchString
	isNumeric       = regexp.MustCompile(`^[0-9]+$`).MatchString
)

// Add a new function to validate the customer_rid
func ValidateCustomerRID(customerRID string) error {
	// Ensure the customer RID contains only numbers and has a length between 6 and 18 (general national ID range)
	if len(customerRID) < 6 || len(customerRID) > 18 {
		return fmt.Errorf("customer RID must be between 6 and 18 characters long")
	}

	if !isNumeric(customerRID) {
		return fmt.Errorf("customer RID must contain only numbers")
	}

	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

// ValidatePhoneNumber validates if a phone number is in a valid format
func ValidatePhoneNumber(value string) error {
	if len(value) < 10 || len(value) > 15 {
		return fmt.Errorf("phone number must contain between 10 to 15 digits")
	}

	if !isPhoneNumber(value) {
		return fmt.Errorf("phone number is not valid")
	}

	return nil
}

func ValidateTransferAmount(value int64, currency string) error {
	if value < 5000 && currency == util.VND {
		return fmt.Errorf("the minimum amount to make a transaction is 5000 VND")
	}
	if value < 1 && currency == util.USD {
		return fmt.Errorf("the minimum amount to make a transaction is 1 USD")
	}
	return nil
}

// ValidatePhoneNumber validates if a phone number is in a valid format
func ValidateAccountNumber(value string) error {
	if len(value) != 11 {
		return fmt.Errorf("account number must contain only 11 digits, you have %d", len(value))
	}

	if !isNumeric(value) {
		return fmt.Errorf("account number must contain only digits")
	}

	return nil
}

// ValidateString ensures the string length is within the allowed range.
func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain between %d and %d characters", minLength, maxLength)
	}
	return nil
}

// ValidateUsername checks if the username is valid (lowercase letters, digits, underscore).
func ValidateUsername(value string) error {
	if err := ValidateString(value, 6, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("username must contain only lowercase letters, digits, or underscore")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

// ValidatePassword checks if the password meets the complexity criteria.
func ValidatePassword(value string) error {
	if len(value) < 8 {
		return fmt.Errorf("password must contain at least 8 characters")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range value {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		case unicode.IsSpace(char):
			return fmt.Errorf("password cannot contain spaces")
		}
	}

	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one digit")
	}
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}
	return nil
}

// ValidateTransactionType checks if the transaction type is valid.
func ValidateTransactionType(value string) error {
	if !util.IsSupportedTransactionType(value) {
		return fmt.Errorf("unsupported transaction type")
	}
	return nil
}

// ValidateCustomerSegment checks if the customer segment is valid.
func ValidateCustomerSegment(value string) error {
	if !util.IsSupportedCustomerSegment(value) {
		return fmt.Errorf("unsupported customer segment")
	}
	return nil
}

// ValidateCustomerTier checks if the customer tier is valid.
func ValidateCustomerTier(value string) error {
	if !util.IsSupportedCustomerTier(value) {
		return fmt.Errorf("unsupported customer tier")
	}
	return nil
}

// ValidateFinancialStatus checks if the financial status is valid.
func ValidateFinancialStatus(value string) error {
	if !util.IsSupportedFinancialStatus(value) {
		return fmt.Errorf("unsupported financial status")
	}
	return nil
}

// ValidateEmployeePosition checks if the employee position is valid.
func ValidateEmployeePosition(value string) error {
	if !util.IsSupportedEmployeePosition(value) {
		return fmt.Errorf("unsupported employee position")
	}
	return nil
}

// ValidateEmployeeStatus checks if the employee status is valid.
func ValidateEmployeeStatus(value string) error {
	if !util.IsSupportedEmployeeStatus(value) {
		return fmt.Errorf("unsupported employee status")
	}
	return nil
}

// ValidateBankStatus checks if the bank status is valid.
func ValidateBankStatus(value string) error {
	if !util.IsSupportedBankStatus(value) {
		return fmt.Errorf("unsupported bank status")
	}
	return nil
}

// ValidateBranchStatus checks if the branch status is valid.
func ValidateBranchStatus(value string) error {
	if !util.IsSupportedBranchStatus(value) {
		return fmt.Errorf("unsupported branch status")
	}
	return nil
}

// ValidateAccountStatus checks if the account status is valid.
func ValidateAccountStatus(value string) error {
	if !util.IsSupportedAccountStatus(value) {
		return fmt.Errorf("unsupported account status")
	}
	return nil
}

// ValidateCurrency checks if the currency type is valid.
func ValidateCurrency(value string) error {
	if !util.IsSupportedCurrencyType(value) {
		return fmt.Errorf("unsupported currency type")
	}
	return nil
}

// ValidateMaturityInstruction checks if the maturity instruction is valid.
func ValidateMaturityInstruction(value string) error {
	if !util.IsSupportedMaturityInstruction(value) {
		return fmt.Errorf("unsupported maturity instruction")
	}
	return nil
}

// ValidateTransactionStatus checks if the transaction status is valid.
func ValidateTransactionStatus(value string) error {
	if !util.IsSupportedTransactionStatus(value) {
		return fmt.Errorf("unsupported transaction status")
	}
	return nil
}

// ValidateSavingStatus checks if the saving status is valid.
func ValidateSavingStatus(value string) error {
	if !util.IsSupportedSavingStatus(value) {
		return fmt.Errorf("unsupported saving status")
	}
	return nil
}

// ValidateLoanStatus checks if the loan status is valid.
func ValidateLoanStatus(value string) error {
	if !util.IsSupportedLoanStatus(value) {
		return fmt.Errorf("unsupported loan status")
	}
	return nil
}

// ValidateDateOfBirth checks if the date of birth can be parsed and is a valid past date.
func ValidateDateOfBirth(dateOfBirth string) error {
	// Define the expected date format
	const layout = "2006/01/02"

	// Try to parse the dateOfBirth using the expected layout
	parsedDate, err := time.Parse(layout, dateOfBirth)
	if err != nil {
		return fmt.Errorf("failed to parse date of birth: %v", err)
	}

	// Check if the parsed date is in the past (valid date of birth)
	if parsedDate.After(time.Now()) {
		return fmt.Errorf("date of birth cannot be in the future")
	}

	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
