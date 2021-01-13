package error

import "fmt"

var (
	// ErrInvalidURL means given url is invalid
	ErrInvalidURL = fmt.Errorf("invalid url")

	// ErrNotFound meas objects is not found
	ErrNotFound = fmt.Errorf("not found")

	// ErrNeedAuth means without correct authentication
	ErrNeedAuth = fmt.Errorf("need correct authentication")

	// ErrNotImpl meas this function not implement yet
	ErrNotImpl = fmt.Errorf("not implement")

	// ErrContentType means content type error
	ErrContentType = fmt.Errorf("content type error")
)

var (
	// ErrInvalidID means given id is invalid
	ErrInvalidID = fmt.Errorf("invalid id")

	// ErrTaskNotExists means the task not exists
	ErrTaskNotExists = fmt.Errorf("task not exists")
)

