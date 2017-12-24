package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewUnit(t *testing.T) {
	u := unit.NewUnit()
	assert.Equal(t, unit.TypeUnit, u.Type())
	assert.NotNil(t, u.ID())
	assert.NotNil(t, u.Created())
	assert.NotNil(t, u.Updated())
}

func TestBaseUnit_ID(t *testing.T) {
	assert.NotEqual(t, unit.NewUnit().ID(), unit.NewUnit().ID())
}

func TestBaseUnit_Title(t *testing.T) {
	u := unit.NewUnit()
	u.SetTitle("MyUnit")
	assert.Equal(t, "MyUnit", u.Title())
}
