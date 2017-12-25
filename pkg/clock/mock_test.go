package clock_test

import (
	"testing"

	"time"

	"github.com/anothermemory/lib/pkg/clock"
	"github.com/stretchr/testify/assert"
)

var dummyTime = time.Date(2017, 11, 24, 17, 0, 0, 0, time.Local)

func TestNewMock(t *testing.T) {
	c := clock.NewMock(dummyTime)
	assert.Equal(t, dummyTime, c.Now())
}

func TestMockClock_Now_SameValue(t *testing.T) {
	c := clock.NewMock(dummyTime)
	assert.Equal(t, c.Now(), c.Now())
}
