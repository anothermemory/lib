package unit

import (
	"testing"

	"encoding/json"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestList_MarshalJSON(t *testing.T) {
	u := newBaseList()
	c1 := NewTextPlain()
	c1.SetTitle("MyText")
	c1.SetData("abc")
	c2 := NewTextCode()
	c2.SetTitle("MyCode")
	c2.SetData("def")
	c2.SetLanguage("PHP")
	u.SetItems([]Unit{c1, c2})

	freezeUnitTime(&u.baseUnit, createdTime, updatedTime)

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)

	u1 := toJSON(jsonUnit(c1), jsonTextPlain(c1.Data()))
	u2 := toJSON(jsonUnit(c2), jsonTextPlain(c2.Data()), jsonTextCode(c2.Language()))

	assert.JSONEq(t, toJSON(jsonUnit(u), fmt.Sprintf(`"items":[%s, %s]`, u1, u2)), string(bytes))
}

func TestList_UnmarshalJSON(t *testing.T) {
	u := NewList()

	u1 := toJSON(jsonUnitDummy(TypeTextPlain), jsonTextPlain("abc"))
	u2 := toJSON(jsonUnitDummy(TypeTextCode), jsonTextPlain("def"), jsonTextCode("PHP"))

	err := json.Unmarshal([]byte(toJSON(jsonUnit(u), fmt.Sprintf(`"items":[%s, %s]`, u1, u2))), &u)
	assert.NoError(t, err)
	items := u.Items()

	tmp := items[0]
	i1, ok := tmp.(TextPlain)
	assert.True(t, ok)
	assert.Equal(t, "abc", i1.Data())

	tmp = items[1]
	i2, ok := tmp.(TextCode)
	assert.True(t, ok)
	assert.Equal(t, "def", i2.Data())
	assert.Equal(t, "PHP", i2.Language())
}
