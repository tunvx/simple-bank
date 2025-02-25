package val

import (
	"fmt"
	"regexp"
	"unicode"
)

func ValidateShardID(shardID int32) error {
	if shardID <= 0 {
		return fmt.Errorf("Shard ID must be greater than 0, starting from 1")
	}
	return nil
}

// Regex patterns for username and other validations
var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isPhoneNumber   = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`).MatchString
	isNumeric       = regexp.MustCompile(`^[0-9]+$`).MatchString
)

// validateStringLength ensures the string length is within the allowed range.
func validateStringLength(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain between %d and %d characters", minLength, maxLength)
	}
	return nil
}

// ValidateUsername checks if the username is valid (lowercase letters, digits, underscore).
func ValidateUsername(value string) error {
	if err := validateStringLength(value, 6, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("username must contain only lowercase letters, digits, or underscore")
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
