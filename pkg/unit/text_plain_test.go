package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextPlain(t *testing.T) {
	u := unit.NewTextPlain("MyUnit", "MyData")
	assert.NotNil(t, u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, "MyData", u.Data())
}

func TestTextPlain_Data(t *testing.T) {
	u := unit.NewTextPlain("MyUnit", "abc")
	u.SetData("MyData")
	assert.Equal(t, "MyData", u.Data())
}
