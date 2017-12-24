package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	u := unit.NewList()
	assert.Equal(t, unit.TypeList, u.Type())
	assert.Len(t, u.Items(), 0)
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
