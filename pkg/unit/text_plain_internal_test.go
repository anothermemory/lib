package unit

import (
	"fmt"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

const formatJSONTextPlain = `"data": "%s"`

func jsonTextPlain(data string) string {
	return fmt.Sprintf(formatJSONTextPlain, data)
}

func TestBaseTextPlain_MarshalJSON(t *testing.T) {
	u := newBaseTextPlain()
	u.SetData("abc")
	u.SetTitle("MyUnit")
	freezeUnitTime(&u.baseUnit, createdTime, updatedTime)

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, toJSON(jsonUnit(u), jsonTextPlain(u.Data())), string(bytes))
}

func TestBaseTextPlain_UnmarshalJSON(t *testing.T) {
	u := NewTextPlain()

	err := json.Unmarshal([]byte(toJSON(jsonUnitDummy(TypeTextPlain), jsonTextPlain("abc"))), &u)
	assert.NoError(t, err)
	assert.Equal(t, "abc", u.Data())
}
