package unit_test

import (
	"testing"

	"encoding/json"
	"fmt"

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

func TestUnit_MarshalJSON(t *testing.T) {
	u := unit.NewUnit("MyUnit")

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, fmt.Sprintf(`{"id": "%s", "title": "MyUnit", "type": "unit"}`, u.ID()), string(bytes))
}

func TestUnit_UnmarshalJSON(t *testing.T) {
	u := unit.NewUnit("")

	err := json.Unmarshal([]byte(`{"id": "123", "title": "MyUnit", "type": "unit"}`), &u)
	assert.NoError(t, err)
	assert.Equal(t, "123", u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, unit.TypeUnit, u.Type())
}

func TestUnit_Type(t *testing.T) {
	assert.Equal(t, unit.TypeUnit, unit.NewUnit("").Type())
}
