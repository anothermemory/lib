package unit_test

import (
	"testing"

	"encoding/json"
	"fmt"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	u := unit.NewList("MyUnit")
	assert.NotNil(t, u.ID())
	assert.NotEmpty(t, u.ID())
	assert.Equal(t, u.Title(), "MyUnit")
	assert.Empty(t, u.Items())
	assert.Len(t, u.Items(), 0)
}

func TestList_AddItem(t *testing.T) {
	u := unit.NewList("MyUnit")
	c1 := unit.NewTextPlain("MyText", "abc")
	c2 := unit.NewTextCode("MyCode", "def", "PHP")

	assert.Empty(t, u.Items())
	assert.Len(t, u.Items(), 0)

	u.AddItem(c1)
	assert.NotEmpty(t, u.Items())
	assert.Len(t, u.Items(), 1)

	u.AddItem(c2)
	assert.Len(t, u.Items(), 2)
}

func TestList_GetItem(t *testing.T) {
	u := unit.NewList("MyUnit")
	c1 := unit.NewTextPlain("MyText", "abc")
	c2 := unit.NewTextCode("MyCode", "def", "PHP")

	u.AddItem(c1)
	u.AddItem(c2)

	tmp := u.GetItem(0)
	i1, ok := tmp.(unit.TextPlain)
	assert.True(t, ok)
	assert.Equal(t, "MyText", i1.Title())
	assert.Equal(t, "abc", i1.Data())

	tmp = u.GetItem(1)
	i2, ok := tmp.(unit.TextCode)
	assert.True(t, ok)
	assert.Equal(t, "MyCode", i2.Title())
	assert.Equal(t, "def", i2.Data())
	assert.Equal(t, "PHP", i2.Language())
}

func TestList_SetItem(t *testing.T) {
	u := unit.NewList("MyUnit")
	c1 := unit.NewTextPlain("MyText", "abc")
	c2 := unit.NewTextCode("MyCode", "def", "PHP")

	u.AddItem(c1)

	tmp := u.GetItem(0)
	i1, ok := tmp.(unit.TextPlain)
	assert.True(t, ok)
	assert.Equal(t, "MyText", i1.Title())
	assert.Equal(t, "abc", i1.Data())

	u.SetItem(0, c2)
	tmp = u.GetItem(0)
	i2, ok := tmp.(unit.TextCode)
	assert.True(t, ok)
	assert.Equal(t, "MyCode", i2.Title())
	assert.Equal(t, "def", i2.Data())
	assert.Equal(t, "PHP", i2.Language())
}

func TestList_RemoveItem(t *testing.T) {
	u := unit.NewList("MyUnit")
	c1 := unit.NewTextPlain("MyText", "abc")
	c2 := unit.NewTextCode("MyCode", "def", "PHP")

	assert.Empty(t, u.Items())
	assert.Len(t, u.Items(), 0)

	u.AddItem(c1)
	u.AddItem(c2)

	u.RemoveItem(0)
	assert.Len(t, u.Items(), 1)
	u.RemoveItem(0)
	assert.Len(t, u.Items(), 0)
}

func TestList_Items(t *testing.T) {
	u := unit.NewList("MyUnit")
	u.SetItems([]unit.Unit{
		unit.NewTextPlain("MyText", "abc"),
		unit.NewTextCode("MyCode", "def", "PHP")})

	items := u.Items()

	tmp := items[0]
	i1, ok := tmp.(unit.TextPlain)
	assert.True(t, ok)
	assert.Equal(t, "MyText", i1.Title())
	assert.Equal(t, "abc", i1.Data())

	tmp = items[1]
	i2, ok := tmp.(unit.TextCode)
	assert.True(t, ok)
	assert.Equal(t, "MyCode", i2.Title())
	assert.Equal(t, "def", i2.Data())
	assert.Equal(t, "PHP", i2.Language())
}

func TestList_MarshalJSON(t *testing.T) {
	u := unit.NewList("MyUnit")
	c1 := unit.NewTextPlain("MyText", "abc")
	c2 := unit.NewTextCode("MyCode", "def", "PHP")
	u.SetItems([]unit.Unit{c1, c2})

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, fmt.Sprintf(`{"id": "%s", "title": "MyUnit", "type":"list", "items":[
{"id": "%s", "title": "MyText", "type":"text_plain", "data":"abc"},
{"id": "%s", "title": "MyCode", "type":"text_code", "data":"def", "language":"PHP"}
]}`, u.ID(), c1.ID(), c2.ID()), string(bytes))
}

func TestList_UnmarshalJSON(t *testing.T) {
	u := unit.NewList("")

	err := json.Unmarshal([]byte(`{"id": "123", "title": "MyUnit", "type":"list", "items":[
{"id": "456", "title": "MyText", "type":"text_plain", "data":"abc"},
{"id": "789", "title": "MyCode", "type":"text_code", "data":"def", "language":"PHP"}
]}`), &u)
	assert.NoError(t, err)
	assert.Equal(t, "123", u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, unit.TypeList, u.Type())

	items := u.Items()

	tmp := items[0]
	i1, ok := tmp.(unit.TextPlain)
	assert.True(t, ok)
	assert.Equal(t, "MyText", i1.Title())
	assert.Equal(t, "abc", i1.Data())

	tmp = items[1]
	i2, ok := tmp.(unit.TextCode)
	assert.True(t, ok)
	assert.Equal(t, "MyCode", i2.Title())
	assert.Equal(t, "def", i2.Data())
	assert.Equal(t, "PHP", i2.Language())
}

func TestList_Type(t *testing.T) {
	assert.Equal(t, unit.TypeList, unit.NewList("").Type())
}
