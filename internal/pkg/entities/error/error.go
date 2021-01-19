package error

import "fmt"

var (
	// ErrCTX means given http request not found contextx
	ErrCTX = fmt.Errorf("transfer contextx failure")

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

	// ErrUpdateTask means update a task is failure
	ErrUpdateTask = fmt.Errorf("update a task is failure")

	// ErrGetTaskByID means get task by id is failure
	ErrGetTaskByID = fmt.Errorf("get task by id is failure")

	// ErrListTasks means list all tasks is failure
	ErrListTasks = fmt.Errorf("list all tasks is failure")

	// ErrDeleteTask means delete a task is failure
	ErrDeleteTask = fmt.Errorf("delete a task is failure")
)

var (
	// ErrCreateObjective means create a objective is failure
	ErrCreateObjective = fmt.Errorf("create a objective is failure")

	// ErrListObjectives means list all objectives is failure
	ErrListObjectives = fmt.Errorf("list all objectives is failure")

	// ErrObjectiveNotExists means objective not exists
	ErrObjectiveNotExists = fmt.Errorf("objective not exists")

	// ErrCountObjective means count all objectives is failure
	ErrCountObjective = fmt.Errorf("count objective is failure")

	// ErrUpdateObj means update a objective is failure
	ErrUpdateObj = fmt.Errorf("update a objective is failure")

	// ErrGetObjByID means get objective by id is failure
	ErrGetObjByID = fmt.Errorf("get objective by id is failure")

	// ErrDeleteObj means delete a objective by id is failure
	ErrDeleteObj = fmt.Errorf("delete a objective by id is failure")
)

var (
	// ErrDBConnect means db connect is failure
	ErrDBConnect = fmt.Errorf("db connect is failure")
)

var (
	// ErrReadiness means readiness is failure
	ErrReadiness = fmt.Errorf("readiness is failure")

	// ErrLiveness means liveness is failure
	ErrLiveness = fmt.Errorf("livenss is failure")
)
