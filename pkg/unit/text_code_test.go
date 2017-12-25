package unit_test

import (
	"testing"

	"encoding/json"
	"fmt"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

const formatJSONTextCode = `"language": "%s"`

func jsonTextCode(language string) string {
	return fmt.Sprintf(formatJSONTextCode, language)
}
func TestNewTextCode(t *testing.T) {
	u := unit.NewTextCode()
	assert.Equal(t, unit.TypeTextCode, u.Type())
}

func TestNewTextCode_Data(t *testing.T) {
	const data = "data"
	u := unit.NewTextCode(unit.TextCodeData(data))
	assert.Equal(t, data, u.Data())
}

func TestNewTextCode_Language(t *testing.T) {
	const language = "language"
	u := unit.NewTextCode(unit.TextCodeLanguage(language))
	assert.Equal(t, language, u.Language())
}

func TestBaseTextCode_Language(t *testing.T) {
	u := unit.NewTextCode()
	u.SetLanguage("Java")
	assert.Equal(t, "Java", u.Language())
}

func TestBaseTextCode_MarshalJSON(t *testing.T) {
	u := unit.NewTextCode(
		unit.ClockMock(createdTime, updatedTime),
		unit.Title("MyUnit"),
		unit.TextPlainData("abc"),
		unit.TextCodeLanguage("PHP"),
	)

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, toJSON(jsonUnit(u), jsonTextPlain(u.Data()), jsonTextCode(u.Language())), string(bytes))
}

func TestBaseTextCode_UnmarshalJSON(t *testing.T) {
	u := unit.NewTextCode()

	err := json.Unmarshal([]byte(toJSON(jsonUnitDummy(unit.TypeTextCode), jsonTextPlain("abc"), jsonTextCode("PHP"))), &u)
	assert.NoError(t, err)
	assert.Equal(t, "PHP", u.Language())
}
