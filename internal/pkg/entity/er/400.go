package er

import "net/http"

var (
	// ErrInvalidID means given id is invalid
	ErrInvalidID = newAPPError(http.StatusBadRequest, 40001, "invalid id")

	// ErrInvalidPage means given page is invalid which MUST be greater than 0
	ErrInvalidPage = newAPPError(http.StatusBadRequest, 40002, "page MUST be greater than 0")

	// ErrInvalidSize means given size is invalid which MUST be greater than 0
	ErrInvalidSize = newAPPError(http.StatusBadRequest, 40003, "size MUST be greater than 0")

	// ErrEmptyTitle means give title of task is empty value
	ErrEmptyTitle = newAPPError(http.StatusBadRequest, 40004, "title must be NOT empty")
)
