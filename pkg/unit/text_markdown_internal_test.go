package unit

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestBaseTextMarkdown_MarshalJSON(t *testing.T) {
	u := newBaseTextMarkdown()
	u.SetData("abc")
	u.SetTitle("MyUnit")
	freezeUnitTime(&u.baseUnit, createdTime, updatedTime)

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, toJSON(jsonUnit(u), jsonTextPlain(u.Data())), string(bytes))
}

func TestBaseTextMarkdown_UnmarshalJSON(t *testing.T) {
	u := NewTextMarkdown()

	err := json.Unmarshal([]byte(toJSON(jsonUnitDummy(TypeTextMarkdown), jsonTextPlain("abc"))), &u)
	assert.NoError(t, err)
	assert.Equal(t, "abc", u.Data())
}
