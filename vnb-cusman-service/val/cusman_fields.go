package val

import (
	"fmt"
	"net/mail"
	"regexp"
	"time"

	"github.com/tunvx/simplebank/common/util"
)

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

//*****************************************************************////*****************************************************************//
//*****************************************************************////*****************************************************************//
//*****************************************************************////*****************************************************************//

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

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters or spaces")
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

// ValidateCurrency checks if the currency type is valid.
func ValidateCurrency(value string) error {
	if !util.IsSupportedCurrencyType(value) {
		return fmt.Errorf("unsupported currency type")
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

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
