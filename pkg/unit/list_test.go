package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	u := unit.NewList("MyUnit")
	assert.NotNil(t, u.ID())
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

	u.AddItem(unit.Unit(c1))
	assert.NotEmpty(t, u.Items())
	assert.Len(t, u.Items(), 1)

	u.AddItem(unit.Unit(c2))
	assert.Len(t, u.Items(), 2)
}

func TestList_GetItem(t *testing.T) {
	u := unit.NewList("MyUnit")
	c1 := unit.NewTextPlain("MyText", "abc")
	c2 := unit.NewTextCode("MyCode", "def", "PHP")

	u.AddItem(unit.Unit(c1))
	u.AddItem(unit.Unit(c2))

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

	u.AddItem(unit.Unit(c1))

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

	u.AddItem(unit.Unit(c1))
	u.AddItem(unit.Unit(c2))

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
