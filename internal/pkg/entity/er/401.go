package er

import "net/http"

var (
	// ErrMissingToken means missing token in header
	ErrMissingToken = newAPPError(http.StatusUnauthorized, 40100, "missing token")

	// ErrAuthHeaderFormat means must provide Authorization header with format `Bearer {token}`
	ErrAuthHeaderFormat = newAPPError(http.StatusUnauthorized, 40101, "Must provide Authorization header with format `Bearer {token}`")

	// ErrIncorrectPassword means incorrect password
	ErrIncorrectPassword = newAPPError(http.StatusUnauthorized, 40102, "incorrect password")

	// ErrExpiredToken means token is expired
	ErrExpiredToken = newAPPError(http.StatusUnauthorized, 40103, "token is expired")
)
