package util

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

// ********** PASS CONVERT UUID TO HTTP AND BACK **********

// ConvertUUIDToString converts a UUID to a hex string
func ConvertUUIDToString(id uuid.UUID) (string, error) {
	return hex.EncodeToString(id[:]), nil
}

// ConvertStringToUUID converts a hex string back to a UUID
func ConvertStringToUUID(id string) (uuid.UUID, error) {
	bytes, err := hex.DecodeString(id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid hex UUID string: %w", err)
	}
	return uuid.FromBytes(bytes)
}

// ConvertAccNumberToInt64 converts a string of number to int64
func ConvertAccNumberToInt64(acc_number string) (int64) {
	account_id, _ := strconv.ParseInt(acc_number, 10, 64)
	return account_id
}