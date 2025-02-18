package util

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBLevelValidationError struct {
	Code    string
	Message string
}

func (e *DBLevelValidationError) Error() string {
	return fmt.Sprintf("%s %s", e.Code, e.Message)
}

const (
	// Database error codes
	UniqueViolation     = "23505"
	ForeignKeyViolation = "23503"

	// Custom error codes
	RecordNotFound    = "20003"
	InsufficientFunds = "20002"
)

var (
	// Database errors
	ErrRecordNotFound  = pgx.ErrNoRows
	ErrUniqueViolation = &pgconn.PgError{
		Code: UniqueViolation,
	}
	ErrForeignKeyViolation = &pgconn.PgError{
		Code: ForeignKeyViolation,
	}

	// Some specific errors to check at the database transaction level
	ErrInsufficientFunds = &DBLevelValidationError{Code: InsufficientFunds, Message: "insufficient funds error"}
)

func ErrorCode(err error) string {
	if errors.Is(err, ErrRecordNotFound) {
		return RecordNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}

	var validateErr *DBLevelValidationError
	if errors.As(err, &validateErr) {
		return validateErr.Code
	}
	return ""
}
