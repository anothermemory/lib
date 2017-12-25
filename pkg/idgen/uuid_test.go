package idgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	generator := NewUUID()
	assert.NotEmpty(t, generator.Generate())
}

func TestGeneratorUUID_Generate(t *testing.T) {
	generator := NewUUID()
	assert.NotEqual(t, generator.Generate(), generator.Generate())
}
