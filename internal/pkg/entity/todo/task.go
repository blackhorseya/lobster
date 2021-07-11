package todo

// Task declare a task basic information
type Task struct {
	// ID describe the unique identify code of task
	ID int64 `json:"id" db:"id"`

	// UserID describe the user id
	UserID int64 `json:"user_id" db:"user_id"`

	// ResultID describe the parent key result's id
	ResultID int64 `json:"result_id" db:"result_id"`

	// Title describe the title of task
	Title string `json:"title" db:"title"`

	// Status describe the status of task
	Status Status `json:"status" db:"status"`

	// CreatedAt describe the task create milliseconds
	CreatedAt int64 `json:"created_at" db:"created_at"`
}
