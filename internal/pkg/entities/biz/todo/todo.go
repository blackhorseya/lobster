package todo

import (
	"fmt"
	"time"
)

// Task declare a task basic information
type Task struct {
	// ID describe the unique identify code of task
	ID string `json:"id" db:"id"`

	// Title describe the title of task
	Title string `json:"title" db:"title"`

	// Completed describe the completed of task
	Completed bool `json:"completed" db:"completed"`

	// CreateAt describe the task create milliseconds
	CreateAt int64 `json:"create_at" db:"create_at"`
}

// ToLineByFormat serve caller to print a line by format
func (t *Task) ToLineByFormat(format string) string {
	return fmt.Sprintf(format, t.ID, t.Title, t.Completed, time.Unix(t.CreateAt/1e9, t.CreateAt%1e9))
}
