package task

import (
	"github.com/blackhorseya/lobster/internal/pkg/utils/timex"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

// Task declare a task basic information
type Task struct {
	// ID describe the unique identify code of task
	ID string `json:"id" db:"id"`

	// ResultID describe the parent key result's id
	ResultID string `json:"result_id" db:"result_id"`

	// Title describe the title of task
	Title string `json:"title" db:"title"`

	// Status describe the status of task
	Status Status `json:"status" db:"status"`

	// Completed describe the completed of task
	Completed bool `json:"completed" db:"completed"`

	// CreateAt describe the task create milliseconds
	CreateAt int64 `json:"create_at" db:"create_at"`
}

// ToLine serve caller to print a string slice
func (t *Task) ToLine() []string {
	return []string{
		t.ID,
		t.ResultID,
		t.Title,
		t.Status.String(),
		timex.Unix(t.CreateAt).Format(timeFormat),
	}
}
