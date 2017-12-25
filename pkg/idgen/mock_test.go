package idgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMock(t *testing.T) {
	generator := NewMock("123")
	assert.Equal(t, "123", generator.Generate())
}

func TestGeneratorMock_Generate_SameValue(t *testing.T) {
	generator := NewMock("123")
	assert.Equal(t, generator.Generate(), generator.Generate())
}
