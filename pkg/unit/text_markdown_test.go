package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextMarkdown(t *testing.T) {
	u := unit.NewTextMarkdown()
	assert.Equal(t, unit.TypeTextMarkdown, u.Type())
}

func TestNewTextMarkdown_Data(t *testing.T) {
	const data = "data"
	u := unit.NewTextMarkdown(unit.OptionTextMarkdownData(data))
	assert.Equal(t, data, u.Data())
}

func TestBaseTextMarkdown_MarshalJSON(t *testing.T) {
	u := unit.NewTextMarkdown(
		unit.OptionClockMock(createdTime, updatedTime),
		unit.OptionTitle("MyUnit"),
		unit.OptionTextPlainData("abc"),
	)

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, toJSON(jsonUnit(u), jsonTextPlain(u.Data())), string(bytes))
}

func TestBaseTextMarkdown_UnmarshalJSON(t *testing.T) {
	u := unit.NewTextMarkdown()

	err := json.Unmarshal([]byte(toJSON(jsonUnitDummy(unit.TypeTextMarkdown), jsonTextPlain("abc"))), &u)
	assert.NoError(t, err)
	assert.Equal(t, "abc", u.Data())
}
