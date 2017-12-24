package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextPlain(t *testing.T) {
	u := unit.NewTextPlain()
	assert.Equal(t, unit.TypeTextPlain, u.Type())
}

func TestTextPlain_Data(t *testing.T) {
	u := unit.NewTextPlain()
	u.SetData("MyData")
	assert.Equal(t, "MyData", u.Data())
}
