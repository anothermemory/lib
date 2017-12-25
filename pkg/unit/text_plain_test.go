package unit_test

import (
	"testing"

	"encoding/json"
	"fmt"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

const formatJSONTextPlain = `"data": "%s"`

func jsonTextPlain(data string) string {
	return fmt.Sprintf(formatJSONTextPlain, data)
}

func TestNewTextPlain(t *testing.T) {
	u := unit.NewTextPlain()
	assert.Equal(t, unit.TypeTextPlain, u.Type())
}

func TestNewTextPlain_Data(t *testing.T) {
	const data = "data"
	u := unit.NewTextPlain(unit.TextPlainData(data))
	assert.Equal(t, data, u.Data())
}

func TestTextPlain_Data(t *testing.T) {
	u := unit.NewTextPlain()
	u.SetData("MyData")
	assert.Equal(t, "MyData", u.Data())
}

func TestBaseTextPlain_MarshalJSON(t *testing.T) {
	u := unit.NewTextPlain(
		unit.ClockMock(createdTime, updatedTime),
		unit.Title("MyUnit"),
		unit.TextPlainData("abc"),
	)

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, toJSON(jsonUnit(u), jsonTextPlain(u.Data())), string(bytes))
}

func TestBaseTextPlain_UnmarshalJSON(t *testing.T) {
	u := unit.NewTextPlain()

	err := json.Unmarshal([]byte(toJSON(jsonUnitDummy(unit.TypeTextPlain), jsonTextPlain("abc"))), &u)
	assert.NoError(t, err)
	assert.Equal(t, "abc", u.Data())
}
