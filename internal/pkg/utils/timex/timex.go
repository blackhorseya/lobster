package timex

import "time"

func Unix(t int64) time.Time {
	return time.Unix(t/1e9, t%1e9)
}
