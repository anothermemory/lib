package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewUnit(t *testing.T) {
	u := unit.NewUnit("MyUnit")
	assert.NotNil(t, u.ID())
	assert.Equal(t, u.Title(), "MyUnit")
}

func TestUnit_ID_Generated(t *testing.T) {
	u := unit.NewUnit("MyUnit")
	assert.NotNil(t, u.ID())
}
func TestUnit_ID_Unique(t *testing.T) {
	assert.NotEqual(t, unit.NewUnit("ABC").ID(), unit.NewUnit("ABC").ID())
}

func TestUnit_Title(t *testing.T) {
	assert.Equal(t, "MyUnit", unit.NewUnit("MyUnit").Title())
}
