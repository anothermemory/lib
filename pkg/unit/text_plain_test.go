package unit_test

import (
	"testing"

	"fmt"

	"encoding/json"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextPlain(t *testing.T) {
	u := unit.NewTextPlain("MyUnit", "MyData")
	assert.NotNil(t, u.ID())
	assert.NotEmpty(t, u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, "MyData", u.Data())
}

func TestTextPlain_Data(t *testing.T) {
	u := unit.NewTextPlain("MyUnit", "abc")
	u.SetData("MyData")
	assert.Equal(t, "MyData", u.Data())
}

func TestTextPlain_MarshalJSON(t *testing.T) {
	u := unit.NewTextPlain("MyUnit", "abc")

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, fmt.Sprintf(`{"id": "%s", "title": "MyUnit", "data": "abc", "type":"text_plain"}`, u.ID()), string(bytes))
}

func TestTextPlain_UnmarshalJSON(t *testing.T) {
	u := unit.NewTextPlain("", "")

	err := json.Unmarshal([]byte(`{"id": "123", "title": "MyUnit", "data": "abc", "type":"text_plain"}`), &u)
	assert.NoError(t, err)
	assert.Equal(t, "123", u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, "abc", u.Data())
	assert.Equal(t, unit.TypeTextPlain, u.Type())
}

func TestTextPlain_Type(t *testing.T) {
	assert.Equal(t, unit.TypeTextPlain, unit.NewTextPlain("", "").Type())
}
