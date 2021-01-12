package todo

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
