package clock

import "time"

type mockClock struct {
	instant time.Time
}

func (c mockClock) Now() time.Time {
	return c.instant
}

// NewMock returns Clock which will return passed time each time
func NewMock(t time.Time) Clock {
	return &mockClock{instant: t}
}
