package okr

// Goal declare an objective basic information
type Goal struct {
	// ID describe the unique identify code of objective
	ID int64 `json:"id" db:"id"`

	// UserID describe the user id
	UserID int64 `json:"user_id" db:"user_id"`

	// Title describe the title of objective
	Title string `json:"title" db:"title"`

	// KeyResults describe key results of objective
	KeyResults []*Result `json:"key_results" db:"key_results"`

	// StartAt describe the objective start timex milliseconds
	StartAt int64 `json:"start_at" db:"start_at"`

	// EndAt describe the objective end timex milliseconds
	EndAt int64 `json:"end_at" db:"end_at"`

	// CreatedAt describe the objective create milliseconds
	CreatedAt int64 `json:"created_at" db:"created_at"`
}
