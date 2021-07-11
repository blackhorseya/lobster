package okr

// Result declare a key result basic information
type Result struct {
	// ID describe the unique identify code of key result
	ID string `json:"id" db:"id"`

	// GoalID describe the parent goal's id
	GoalID string `json:"goal_id" db:"goal_id"`

	// Title describe the title of key result
	Title string `json:"title" db:"title"`

	// Target describe the target of key result
	Target int `json:"target" db:"target"`

	// Actual describe the actual of key result
	Actual int `json:"actual" db:"actual"`

	// Progress describe the progress of key result
	Progress int `json:"-"`

	// CreatedAt describe the key result create milliseconds
	CreatedAt int64 `json:"created_at" db:"created_at"`
}
