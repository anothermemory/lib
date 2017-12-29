package unit_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	u := unit.NewList()
	assert.Equal(t, unit.TypeList, u.Type())
	assert.Len(t, u.Items(), 0)
}

func TestNewList_Item(t *testing.T) {
	u := unit.NewList(
		unit.OptionListItem(
			unit.NewTextPlain(unit.OptionTitle("Text1"), unit.OptionTextPlainData("Data1")),
		),
		unit.OptionListItem(
			unit.NewTextPlain(unit.OptionTitle("Text2"), unit.OptionTextPlainData("Data1")),
		),
	)

	assert.Equal(t, "Text1", u.Items()[0].Title())
	assert.Equal(t, "Text2", u.Items()[1].Title())
}

func TestList_AddItem(t *testing.T) {
	u := unit.NewList()

	assert.Empty(t, u.Items())
	assert.Len(t, u.Items(), 0)

	u.AddItem(unit.NewTextPlain())
	assert.NotEmpty(t, u.Items())
	assert.Len(t, u.Items(), 1)

	u.AddItem(unit.NewTextCode())
	assert.Len(t, u.Items(), 2)
}

func TestList_GetItem(t *testing.T) {
	u := unit.NewList()
	c1 := unit.NewTextPlain()
	c1.SetTitle("MyText")
	c1.SetData("abc")
	c2 := unit.NewTextCode()
	c2.SetTitle("MyCode")
	c2.SetData("def")
	c2.SetLanguage("PHP")

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
	u := unit.NewList()
	c1 := unit.NewTextPlain()
	c1.SetTitle("MyText")
	c1.SetData("abc")
	c2 := unit.NewTextCode()
	c2.SetTitle("MyCode")
	c2.SetData("def")
	c2.SetLanguage("PHP")

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
	u := unit.NewList()
	c1 := unit.NewTextPlain()
	c1.SetTitle("MyText")
	c1.SetData("abc")
	c2 := unit.NewTextCode()
	c2.SetTitle("MyCode")
	c2.SetData("def")
	c2.SetLanguage("PHP")

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
	u := unit.NewList()
	c1 := unit.NewTextPlain()
	c1.SetTitle("MyText")
	c1.SetData("abc")
	c2 := unit.NewTextCode()
	c2.SetTitle("MyCode")
	c2.SetData("def")
	c2.SetLanguage("PHP")
	u.SetItems([]unit.Unit{c1, c2})

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
	u := unit.NewList(
		unit.OptionClockMock(createdTime, updatedTime),
		unit.OptionTitle("MyUnit"),
	)
	c1 := unit.NewTextPlain()
	c1.SetTitle("MyText")
	c1.SetData("abc")
	c2 := unit.NewTextCode()
	c2.SetTitle("MyCode")
	c2.SetData("def")
	c2.SetLanguage("PHP")
	u.SetItems([]unit.Unit{c1, c2})

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)

	u1 := toJSON(jsonUnit(c1), jsonTextPlain(c1.Data()))
	u2 := toJSON(jsonUnit(c2), jsonTextPlain(c2.Data()), jsonTextCode(c2.Language()))

	assert.JSONEq(t, toJSON(jsonUnit(u), fmt.Sprintf(`"items":[%s, %s]`, u1, u2)), string(bytes))
}

func TestList_UnmarshalJSON(t *testing.T) {
	u := unit.NewList()

	u1 := toJSON(jsonUnitDummy(unit.TypeTextPlain), jsonTextPlain("abc"))
	u2 := toJSON(jsonUnitDummy(unit.TypeTextCode), jsonTextPlain("def"), jsonTextCode("PHP"))

	err := json.Unmarshal([]byte(toJSON(jsonUnit(u), fmt.Sprintf(`"items":[%s, %s]`, u1, u2))), &u)
	assert.NoError(t, err)
	items := u.Items()

	tmp := items[0]
	i1, ok := tmp.(unit.TextPlain)
	assert.True(t, ok)
	assert.Equal(t, "abc", i1.Data())

	tmp = items[1]
	i2, ok := tmp.(unit.TextCode)
	assert.True(t, ok)
	assert.Equal(t, "def", i2.Data())
	assert.Equal(t, "PHP", i2.Language())
}

func TestList_MarshalJSON_WithoutItems(t *testing.T) {
	u := unit.NewList(unit.ListMarshalItems(false))
	c1 := unit.NewTextPlain()
	c1.SetTitle("MyText")
	c1.SetData("abc")
	c2 := unit.NewTextPlain()
	c2.SetTitle("MyCode")
	c2.SetData("def")
	u.SetItems([]unit.Unit{c1, c2})

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)

	assert.JSONEq(t, toJSON(jsonUnit(u), fmt.Sprintf(`"items":["%s", "%s"]`, c1.ID(), c2.ID())), string(bytes))
}

func TestList_UnmarshalJSON_WithoutItems(t *testing.T) {
	u := unit.NewList(unit.ListMarshalItems(false))
	const i1ID = "123"
	const i2ID = "456"

	err := json.Unmarshal([]byte(toJSON(jsonUnit(u), fmt.Sprintf(`"items":["%s", "%s"]`, i1ID, i2ID))), &u)
	assert.NoError(t, err)
	items := u.Items()

	tmp := items[0]
	assert.Equal(t, unit.TypeUnit, tmp.Type())
	i1, ok := tmp.(unit.Unit)
	assert.True(t, ok)
	assert.Equal(t, i1ID, i1.ID())

	tmp = items[1]
	assert.Equal(t, unit.TypeUnit, tmp.Type())
	i2, ok := tmp.(unit.Unit)
	assert.True(t, ok)
	assert.Equal(t, i2ID, i2.ID())
}
