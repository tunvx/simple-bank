package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

var rnd *rand.Rand

func init() {
	// Create a new random generator with a seed based on the current time
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomCustomerSegment generates a random customer segment
func RandomCustomerSegment() string {
	segments := []string{
		"individual",
		"small_enterprise",
		"medium_enterprise",
		"large_enterprise",
		"institutional",
	}

	return segments[rnd.Intn(len(segments))]
}

func RandomCustomerTier() string {
	tiers := []string{
		"regular",
		"bronze",
		"silver",
		"gold",
		"platinum",
		"diamond",
	}

	n := len(tiers)
	return tiers[rnd.Intn(n)]
}

func RandomFinancialStatus() string {
	statuses := []string{
		"excellent",
		"very_good",
		"good",
		"fair",
		"poor",
		"very_poor",
	}

	n := len(statuses)
	return statuses[rnd.Intn(n)]
}

func RandomEmployeePosition() string {
	positions := []string{
		"financial_analyst",
		"banker",
		"branch_manager",
	}

	n := len(positions)
	return positions[rnd.Intn(n)]
}

func RandomEmployeeStatus() string {
	statuses := []string{
		"active",
		"inactive",
	}

	n := len(statuses)
	return statuses[rnd.Intn(n)]
}

func RandomBankStatus() string {
	statuses := []string{
		"active",
		"inactive",
	}

	n := len(statuses)
	return statuses[rnd.Intn(n)]
}

func RandomBranchStatus() string {
	statuses := []string{
		"active",
		"inactive",
	}

	n := len(statuses)
	return statuses[rnd.Intn(n)]
}

func RandomAccountStatus() string {
	statuses := []string{
		"active",
		"inactive",
	}

	n := len(statuses)
	return statuses[rnd.Intn(n)]
}

func RandomCurrencyType() string {
	types := []string{
		"VND",
		"USD",
	}

	n := len(types)
	return types[rnd.Intn(n)]
}

func RandomMaturityInstruction() string {
	instructions := []string{
		"reinvest",
		"interest_only",
		"full_withdrawal",
	}

	n := len(instructions)
	return instructions[rnd.Intn(n)]
}

func RandomTransactionStatus() string {
	statuses := []string{
		"pending",
		"completed",
		"failed",
	}

	n := len(statuses)
	return statuses[rnd.Intn(n)]
}

func RandomSavingStatus() string {
	statuses := []string{
		"active",
		"completed",
		"terminated",
		"inactive",
	}

	n := len(statuses)
	return statuses[rnd.Intn(n)]
}

func RandomLoanStatus() string {
	statuses := []string{
		"accruing",
		"paid_off",
		"extended",
		"non_performing",
	}

	n := len(statuses)
	return statuses[rnd.Intn(n)]
}

func RandomPersionID() string {
	return gofakeit.CreditCardNumber(nil)[:12]
}

func RandomPersonName() string {
	return gofakeit.Name()
}

func RandomDate() time.Time {
	randomDate := gofakeit.Date()
	year, month, day := randomDate.Date()
	// Create a new time.Time object with only the year, month, and day
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func DateToString(date time.Time) string {
	year, month, day := date.Date()
	// Create a new time.Time object with only the year, month, and day
	date = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	// Return the date as a string in the desired format
	return date.Format("2006/01/02")
}

func RandomAddress() string {
	return gofakeit.Address().Address
}

func RandomPhoneNumber() string {
	return gofakeit.Phone()
}

func RandomEmail() string {
	return gofakeit.Email()
}

func RandomBranchCode() string {
	return gofakeit.Regex("BR[0-9]{5}")
}

func RandomBranchName() string {
	return gofakeit.Company()
}

func RandomBankCode() string {
	return gofakeit.Regex("BR[0-9]{5}")
}

func RandomBankName() string {
	return gofakeit.Company()
}

func RandomAccountNumber() string {
	return gofakeit.CreditCardNumber(nil)[:11]
}

func RandomEmployeeCode() string {
	return gofakeit.CreditCardNumber(nil)[:8]
}

func RandomLoanCode() string {
	return gofakeit.CreditCardNumber(nil)[:11]
}

func RandomSavingCode() string {
	return gofakeit.CreditCardNumber(nil)[:11]
}

func RandomLoanName() string {
	return gofakeit.Product().Name
}

func RandomSavingName() string {
	return gofakeit.Product().Name
}

func RandomInt64ID() int64 {
	for {
		id := gofakeit.Int64()
		if id > 0 {
			return id
		}
	}
}

func RandomInt16ID() int16 {
	for {
		id := gofakeit.Int16()
		if id > 0 {
			return id
		}
	}
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rnd.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

const (
	passwordLength    = 12
	lowercaseLetters  = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers           = "0123456789"
	specialCharacters = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
	allCharacters     = lowercaseLetters + uppercaseLetters + numbers + specialCharacters
)

// RandomString generates a random string of length n
func RandomUsername() string {
	allCharacters := lowercaseLetters + numbers
	minLength := 8
	maxLength := 15

	length := rnd.Intn(maxLength-minLength+1) + minLength

	username := make([]byte, length)
	for i := range username {
		username[i] = allCharacters[rnd.Intn(len(allCharacters))]
	}

	return string(username)
}

func randomCharacterFromSet(set string) rune {
	return rune(set[rand.Intn(len(set))])
}

func RandomPassword() (string, error) {
	var password strings.Builder

	// Ensure the password has at least one of each required character type
	password.WriteRune(randomCharacterFromSet(lowercaseLetters))
	password.WriteRune(randomCharacterFromSet(uppercaseLetters))
	password.WriteRune(randomCharacterFromSet(numbers))
	password.WriteRune(randomCharacterFromSet(specialCharacters))

	// Fill the remaining length of the password with random characters
	for password.Len() < passwordLength {
		password.WriteRune(randomCharacterFromSet(allCharacters))
	}

	// Shuffle the characters in the password
	runes := []rune(password.String())
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})

	finalPassword := string(runes)

	// Validate the generated password
	if err := ValidatePassword(finalPassword); err != nil {
		return "", err
	}

	return finalPassword, nil
}
