package clock_test

import (
	"testing"
	"time"

	"github.com/anothermemory/lib/pkg/clock"
	"github.com/stretchr/testify/assert"
)

func TestNewMockPartial(t *testing.T) {
	onceTime := time.Date(2017, 11, 24, 17, 0, 0, 0, time.Local)
	c := clock.NewMockPartial(onceTime)
	assert.Equal(t, onceTime, c.Now())
	assert.NotEqual(t, onceTime, c.Now())
}

func TestNewMockPartial_MultipleValues(t *testing.T) {
	firstTime := time.Date(2017, 11, 24, 17, 0, 0, 0, time.Local)
	secondTime := time.Date(2018, 11, 24, 17, 0, 0, 0, time.Local)
	c := clock.NewMockPartial(firstTime, secondTime)
	assert.Equal(t, firstTime, c.Now())
	assert.Equal(t, secondTime, c.Now())
	nextTime := c.Now()
	assert.NotEqual(t, firstTime, nextTime)
	assert.NotEqual(t, secondTime, nextTime)
}
