package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextMarkdown(t *testing.T) {
	u := unit.NewTextMarkdown()
	assert.Equal(t, unit.TypeTextMarkdown, u.Type())
}
