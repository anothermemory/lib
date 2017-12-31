package clock

import "time"

type mockPartialClock struct {
	instants []time.Time
	index    int
	useReal  bool
}

func (c *mockPartialClock) Now() time.Time {
	if c.useReal {
		return time.Now()
	}

	defer c.next()
	return c.instants[c.index]
}

func (c *mockPartialClock) next() {
	c.index++
	if c.index == len(c.instants) {
		c.useReal = true
	}
}

// NewMockPartial returns Clock which will return passed time once and then real time each time
func NewMockPartial(t ...time.Time) Clock {
	return &mockPartialClock{instants: t, useReal: false}
}
