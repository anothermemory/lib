package directory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLocation(t *testing.T) {
	l := newLocation("/var/app", "123")
	assert.NotEmpty(t, l.filename)
	assert.NotEmpty(t, l.dirPath)
	assert.NotEmpty(t, l.fullPath)
}

func TestNewLocation_SameID(t *testing.T) {
	l1 := newLocation("/var/app", "123")
	l2 := newLocation("/var/app", "123")
	assert.Equal(t, l1.filename, l2.filename)
	assert.Equal(t, l1.dirPath, l2.dirPath)
	assert.Equal(t, l1.fullPath, l2.fullPath)
}
