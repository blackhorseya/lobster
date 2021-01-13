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

	// ErrInvalidPage means given page is invalid which MUST be greater than 0
	ErrInvalidPage = fmt.Errorf("page MUST be greater than 0")

	// ErrInvalidSize means given size is invalid which MUST be greater than 0
	ErrInvalidSize = fmt.Errorf("size MUST be greater than 0")

	// ErrEmptyTitle means give title of task is empty value
	ErrEmptyTitle = fmt.Errorf("title must be NOT empty")

	// ErrCreateTask means create a task is failure
	ErrCreateTask = fmt.Errorf("create a task is failure")
)
