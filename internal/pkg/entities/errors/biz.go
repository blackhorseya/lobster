package errors

import (
	"net/http"
)

var (
	// ErrReadiness means readiness is failure
	ErrReadiness = newAPPError(http.StatusInternalServerError, 50010, "readiness is failure")

	// ErrLiveness means liveness is failure
	ErrLiveness = newAPPError(http.StatusInternalServerError, 50011, "liveness is failure")
)

var (
	// ErrCreateObj means create a objective is failure
	ErrCreateObj = newAPPError(http.StatusInternalServerError, 50020, "create an objective is failure")

	// ErrListObj means list all objectives is failure
	ErrListObj = newAPPError(http.StatusInternalServerError, 50021, "list all objectives is failure")

	// ErrObjNotExists means objective not exists
	ErrObjNotExists = newAPPError(http.StatusInternalServerError, 50022, "objective not exists")

	// ErrCountObj means count all objectives is failure
	ErrCountObj = newAPPError(http.StatusInternalServerError, 50023, "count objective is failure")

	// ErrUpdateObj means update a objective is failure
	ErrUpdateObj = newAPPError(http.StatusInternalServerError, 50024, "update a objective is failure")

	// ErrGetObjByID means get objective by id is failure
	ErrGetObjByID = newAPPError(http.StatusInternalServerError, 50025, "get objective by id is failure")

	// ErrDeleteObj means delete a objective by id is failure
	ErrDeleteObj = newAPPError(http.StatusInternalServerError, 50026, "delete a objective by id is failure")
)

var (
	// ErrCreateTask means create a task is failure
	ErrCreateTask = newAPPError(http.StatusInternalServerError, 50030, "create a task is failure")

	// ErrUpdateTask means update a task is failure
	ErrUpdateTask = newAPPError(http.StatusInternalServerError, 50031, "update a task is failure")

	// ErrGetTaskByID means get task by id is failure
	ErrGetTaskByID = newAPPError(http.StatusInternalServerError, 50032, "get task by id is failure")

	// ErrListTasks means list all tasks is failure
	ErrListTasks = newAPPError(http.StatusInternalServerError, 50033, "list all tasks is failure")

	// ErrDeleteTask means delete a task is failure
	ErrDeleteTask = newAPPError(http.StatusInternalServerError, 50034, "delete a task is failure")

	// ErrTaskNotExists means task is not exists
	ErrTaskNotExists = newAPPError(http.StatusNotFound, 50035, "task is not exists")
)

var (
	// ErrGetKRByID means get a key result by id is failure
	ErrGetKRByID = newAPPError(http.StatusInternalServerError, 50040, "get a key result by id is failure")

	// ErrKRNotExists means key result not exists
	ErrKRNotExists = newAPPError(http.StatusNotFound, 50041, "key result not exists")

	// ErrListKR means list all key results is failure
	ErrListKR = newAPPError(http.StatusInternalServerError, 50042, "list all key results is failure")

	// ErrDeleteKR means delete a key result is failure
	ErrDeleteKR = newAPPError(http.StatusInternalServerError, 50043, "delete a key result is failure")

	// ErrUpdateKR means update a key result is failure
	ErrUpdateKR = newAPPError(http.StatusInternalServerError, 50044, "update a key result is failure")

	// ErrCreateKR means create a key result is failure
	ErrCreateKR = newAPPError(http.StatusInternalServerError, 50045, "create a key result is failure")
)
