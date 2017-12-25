package clock

import "time"

// Clock represents interface which can be used to get current time
type Clock interface {
	Now() time.Time
}

type realClock struct {
}

func (realClock) Now() time.Time {
	return time.Now()
}

// NewReal returns Clock which will return actual real time
func NewReal() Clock {
	return &realClock{}
}
