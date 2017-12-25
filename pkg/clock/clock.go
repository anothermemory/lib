package clock

import "time"

// https://github.com/smartystreets/clock/blob/master/clock.go
type Clock struct {
	instant time.Time
}

func (c *Clock) Now() time.Time {
	if c == nil {
		return time.Now()
	}

	return c.instant
}
func Freeze(instant time.Time) *Clock {
	return &Clock{instant: instant}
}
