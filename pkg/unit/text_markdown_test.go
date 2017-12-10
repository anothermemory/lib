package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextMarkdown(t *testing.T) {
	u := unit.NewTextMarkdown("MyUnit", "MyData")
	assert.NotNil(t, u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, "MyData", u.Data())
}
