package okr

import (
	"fmt"
	"time"

)

// Objective declare an objective basic information
type Objective struct {
	// ID describe the unique identify code of objective
	ID string `json:"id" db:"id"`

	// Title describe the title of objective
	Title string `json:"title" db:"title"`

	// KeyResults describe key results of objective
	KeyResults []*KeyResult `json:"key_results" db:"key_results"`

	// StartAt describe the objective start time milliseconds
	StartAt int64 `json:"start_at" db:"start_at"`

	// EndAt describe the objective end time milliseconds
	EndAt int64 `json:"end_at" db:"end_at"`

	// CreateAt describe the objective create milliseconds
	CreateAt int64 `json:"create_at" db:"create_at"`
}

func (o *Objective) ToLineByFormat(format string) string {
	return fmt.Sprintf(format, o.ID, o.Title, time.Unix(o.CreateAt/1e9, o.CreateAt%1e9))
}

// KeyResult declare a key result basic information
type KeyResult struct {
	// ID describe the unique identify code of key result
	ID string `json:"id" db:"id"`

	// Title describe the title of key result
	Title string `json:"title" db:"title"`

	// Target describe the target of key result
	Target int64 `json:"target" db:"target"`

	// Actual describe the actual of key result
	Actual int64 `json:"actual" db:"actual"`

	// CreateAt describe the key result create milliseconds
	CreateAt int64 `json:"create_at" db:"create_at"`
}
