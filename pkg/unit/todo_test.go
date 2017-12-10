package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTodoItem(t *testing.T) {
	i := unit.NewTodoItem("abc", true)
	assert.Equal(t, "abc", i.Data())
	assert.True(t, i.Done())
}

func TestTodoItem_Data(t *testing.T) {
	u := unit.NewTodoItem("abc", true)
	u.SetData("MyData")
	assert.Equal(t, "MyData", u.Data())
}

func TestTodoItem_Done(t *testing.T) {
	u := unit.NewTodoItem("abc", true)
	u.SetDone(false)
	assert.False(t, u.Done())
}

func TestNewTodo(t *testing.T) {
	u := unit.NewTodo("MyUnit")
	assert.NotNil(t, u.ID())
	assert.Equal(t, u.Title(), "MyUnit")
	assert.Empty(t, u.Items())
	assert.Len(t, u.Items(), 0)
}

func TestTodo_AddItem(t *testing.T) {
	u := unit.NewTodo("MyUnit")
	c1 := unit.NewTodoItem("abc", true)
	c2 := unit.NewTodoItem("def", false)

	assert.Empty(t, u.Items())
	assert.Len(t, u.Items(), 0)

	u.AddItem(c1)
	assert.NotEmpty(t, u.Items())
	assert.Len(t, u.Items(), 1)

	u.AddItem(c2)
	assert.Len(t, u.Items(), 2)
}

func TestTodo_GetItem(t *testing.T) {
	u := unit.NewTodo("MyUnit")
	c1 := unit.NewTodoItem("abc", true)
	c2 := unit.NewTodoItem("def", false)

	u.AddItem(c1)
	u.AddItem(c2)

	tmp := u.GetItem(0)
	assert.Equal(t, "abc", tmp.Data())
	assert.Equal(t, true, tmp.Done())

	tmp = u.GetItem(1)
	assert.Equal(t, "def", tmp.Data())
	assert.Equal(t, false, tmp.Done())
}

func TestTodo_SetItem(t *testing.T) {
	u := unit.NewTodo("MyUnit")
	c1 := unit.NewTodoItem("abc", true)
	c2 := unit.NewTodoItem("def", false)

	u.AddItem(c1)

	tmp := u.GetItem(0)
	assert.Equal(t, "abc", tmp.Data())
	assert.Equal(t, true, tmp.Done())

	u.SetItem(0, c2)
	tmp = u.GetItem(0)
	assert.Equal(t, "def", tmp.Data())
	assert.Equal(t, false, tmp.Done())
}

func TestTodo_RemoveItem(t *testing.T) {
	u := unit.NewTodo("MyUnit")
	c1 := unit.NewTodoItem("abc", true)
	c2 := unit.NewTodoItem("def", false)

	assert.Empty(t, u.Items())
	assert.Len(t, u.Items(), 0)

	u.AddItem(c1)
	u.AddItem(c2)

	u.RemoveItem(0)
	assert.Len(t, u.Items(), 1)
	u.RemoveItem(0)
	assert.Len(t, u.Items(), 0)
}

func TestTodo_Items(t *testing.T) {
	u := unit.NewTodo("MyUnit")
	u.SetItems([]unit.TodoItem{
		unit.NewTodoItem("abc", true),
		unit.NewTodoItem("def", false)})

	items := u.Items()

	tmp := items[0]
	assert.Equal(t, "abc", tmp.Data())
	assert.Equal(t, true, tmp.Done())

	tmp = items[1]
	assert.Equal(t, "def", tmp.Data())
	assert.Equal(t, false, tmp.Done())
}