package er

import "net/http"

var (
	// ErrGetUserByEmail means get user by email is failure
	ErrGetUserByEmail = newAPPError(http.StatusInternalServerError, 50001, "get user by email is failure")

	// ErrEncryptPassword means encrypt password is failure
	ErrEncryptPassword = newAPPError(http.StatusInternalServerError, 50002, "encrypt password is failure")

	// ErrSignup means signup is failure
	ErrSignup = newAPPError(http.StatusInternalServerError, 50003, "signup is failure")

	// ErrNewToken means new a jwt is failure
	ErrNewToken = newAPPError(http.StatusInternalServerError, 50004, "new a jwt is failure")

	// ErrLogin means login is failure
	ErrLogin = newAPPError(http.StatusInternalServerError, 50005, "login is failure")

	// ErrValidateToken means couldn't parse claims
	ErrValidateToken = newAPPError(http.StatusInternalServerError, 50006, "Couldn't parse claims")
)
