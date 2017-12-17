package unit_test

import (
	"testing"

	"encoding/json"
	"fmt"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextCode(t *testing.T) {
	u := unit.NewTextCode("MyUnit", "MyData", "PHP")
	assert.NotNil(t, u.ID())
	assert.NotEmpty(t, u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, "MyData", u.Data())
	assert.Equal(t, "PHP", u.Language())
}

func TestTextCode_Language(t *testing.T) {
	u := unit.NewTextCode("MyUnit", "MyData", "PHP")
	u.SetLanguage("Java")
	assert.Equal(t, "Java", u.Language())
}

func TestTextCode_MarshalJSON(t *testing.T) {
	u := unit.NewTextCode("MyUnit", "abc", "PHP")

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, fmt.Sprintf(`{"id": "%s", "title": "MyUnit", "data": "abc", "language": "PHP", "type":"text_code"}`, u.ID()), string(bytes))
}

func TestTextCode_UnmarshalJSON(t *testing.T) {
	u := unit.NewTextCode("", "", "")

	err := json.Unmarshal([]byte(`{"id": "123", "title": "MyUnit", "data": "abc", "language": "PHP", "type":"text_code"}`), &u)
	assert.NoError(t, err)
	assert.Equal(t, "123", u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, "abc", u.Data())
	assert.Equal(t, "PHP", u.Language())
	assert.Equal(t, "text_code", u.Type())
}

func TestTextCode_Type(t *testing.T) {
	assert.Equal(t, "text_code", unit.NewTextCode("", "", "").Type())
}
