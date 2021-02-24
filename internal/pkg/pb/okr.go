package pb

import (
	"fmt"
	"strconv"
	"time"
)

// Goal declare an objective basic information
type Goal struct {
	// ID describe the unique identify code of objective
	ID string `json:"id" db:"id"`

	// Title describe the title of objective
	Title string `json:"title" db:"title"`

	// KeyResults describe key results of objective
	KeyResults []*KeyResult `json:"key_results" db:"key_results"`

	// StartAt describe the objective start timex milliseconds
	StartAt int64 `json:"start_at" db:"start_at"`

	// EndAt describe the objective end timex milliseconds
	EndAt int64 `json:"end_at" db:"end_at"`

	// CreateAt describe the objective create milliseconds
	CreateAt int64 `json:"create_at" db:"create_at"`
}

// ToLine serve caller to print a string slice
func (o *Goal) ToLine() []string {
	return []string{
		o.ID,
		o.Title,
		time.Unix(o.StartAt/1e9, o.StartAt%1e9).Format(timeFormat),
		time.Unix(o.EndAt/1e9, o.EndAt%1e9).Format(timeFormat),
		time.Unix(o.CreateAt/1e9, o.CreateAt%1e9).Format(timeFormat),
	}
}

// KeyResult declare a key result basic information
type KeyResult struct {
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

	// CreateAt describe the key result create milliseconds
	CreateAt int64 `json:"create_at" db:"create_at"`
}

// ToLine serve caller to print a string slice
func (k *KeyResult) ToLine() []string {
	return []string{
		k.ID,
		k.GoalID,
		k.Title,
		strconv.Itoa(k.Target),
		strconv.Itoa(k.Actual),
		fmt.Sprintf("%.2f", (float32(k.Actual)/float32(k.Target))*100),
		time.Unix(k.CreateAt/1e9, k.CreateAt%1e9).Format(timeFormat),
	}
}
