package er

import "net/http"

var (
	// ErrUserNotExists means user is not exists
	ErrUserNotExists = newAPPError(http.StatusNotFound, 40400, "user is not exists")
)
