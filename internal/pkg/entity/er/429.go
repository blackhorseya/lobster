package er

import "net/http"

var (
	// ErrRateLimit means too many requests
	ErrRateLimit = newAPPError(http.StatusTooManyRequests, 42900, "too many requests")
)
