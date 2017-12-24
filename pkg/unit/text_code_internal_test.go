package unit

import (
	"testing"

	"encoding/json"
	"fmt"

	"github.com/stretchr/testify/assert"
)

const formatJSONTextCode = `"language": "%s"`

func jsonTextCode(language string) string {
	return fmt.Sprintf(formatJSONTextCode, language)
}

func TestBaseTextCode_MarshalJSON(t *testing.T) {
	u := newBaseTextCode()
	u.SetData("abc")
	u.SetTitle("MyUnit")
	u.SetLanguage("PHP")
	freezeUnitTime(&u.baseUnit, createdTime, updatedTime)

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, toJSON(jsonUnit(u), jsonTextPlain(u.Data()), jsonTextCode(u.Language())), string(bytes))
}

func TestBaseTextCode_UnmarshalJSON(t *testing.T) {
	u := NewTextCode()

	err := json.Unmarshal([]byte(toJSON(jsonUnitDummy(TypeTextCode), jsonTextPlain("abc"), jsonTextCode("PHP"))), &u)
	assert.NoError(t, err)
	assert.Equal(t, "PHP", u.Language())
}
