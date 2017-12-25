package clock

import "time"

type mockClock struct {
	instants []time.Time
	index    int
}

func (c *mockClock) Now() time.Time {
	defer c.next()
	return c.instants[c.index]
}

func (c *mockClock) next() {
	c.index++
	if c.index == len(c.instants) {
		c.index = 0
	}
}

// NewMock returns Clock which will return passed time each time
func NewMock(t ...time.Time) Clock {
	return &mockClock{instants: t, index: 0}
}
