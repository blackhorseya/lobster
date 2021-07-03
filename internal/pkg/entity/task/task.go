package task

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

	// CreatedAt describe the task create milliseconds
	CreatedAt int64 `json:"created_at" db:"created_at"`
}
