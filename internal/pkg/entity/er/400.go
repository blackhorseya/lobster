package er

import "net/http"

var (
	// ErrMissingEmail means email is empty
	ErrMissingEmail = newAPPError(http.StatusBadRequest, 40001, "email is empty")

	// ErrMissingPassword means password is empty
	ErrMissingPassword = newAPPError(http.StatusBadRequest, 40002, "password is empty")

	// ErrMissingPath means missing path
	ErrMissingPath = newAPPError(http.StatusBadRequest, 40003, "missing path")

	// ErrMissingDate means missing date
	ErrMissingDate = newAPPError(http.StatusBadRequest, 40004, "missing date")

	// ErrEmailExists means email is exists
	ErrEmailExists = newAPPError(http.StatusBadRequest, 40005, "email is exists")
)
