package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTodoItem(t *testing.T) {
	i := unit.NewTodo().NewItem()
	assert.Implements(t, (*unit.TodoItem)(nil), i)
}

func TestBaseTodoItem_Data(t *testing.T) {
	i := unit.NewTodo().NewItem()
	i.SetData("MyData")
	assert.Equal(t, "MyData", i.Data())
}

func TestBaseTodoItem_Done(t *testing.T) {
	i := unit.NewTodo().NewItem()
	i.SetDone(false)
	assert.False(t, i.Done())
}

func TestNewTodo(t *testing.T) {
	u := unit.NewTodo()
	assert.Equal(t, unit.TypeTodo, u.Type())
}

func TestNewTodo_Item(t *testing.T) {
	u := unit.NewTodo(unit.OptionTodoItem("data1", true), unit.OptionTodoItem("data2", false))
	items := u.Items()

	assert.Equal(t, "data1", items[0].Data())
	assert.Equal(t, true, items[0].Done())

	assert.Equal(t, "data2", items[1].Data())
	assert.Equal(t, false, items[1].Done())
}

func TestTodo_AddItem(t *testing.T) {
	u := unit.NewTodo()
	c1 := u.NewItem()
	c1.SetData("abc")
	c1.SetDone(true)
	c2 := u.NewItem()
	c2.SetData("def")
	c2.SetDone(false)

	assert.Empty(t, u.Items())
	assert.Len(t, u.Items(), 0)

	u.AddItem(c1)
	assert.NotEmpty(t, u.Items())
	assert.Len(t, u.Items(), 1)

	u.AddItem(c2)
	assert.Len(t, u.Items(), 2)
}

func TestTodo_GetItem(t *testing.T) {
	u := unit.NewTodo()
	c1 := u.NewItem()
	c1.SetData("abc")
	c1.SetDone(true)
	c2 := u.NewItem()
	c2.SetData("def")
	c2.SetDone(false)

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
	u := unit.NewTodo()
	c1 := u.NewItem()
	c1.SetData("abc")
	c1.SetDone(true)
	c2 := u.NewItem()
	c2.SetData("def")
	c2.SetDone(false)

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
	u := unit.NewTodo()
	c1 := u.NewItem()
	c1.SetData("abc")
	c1.SetDone(true)
	c2 := u.NewItem()
	c2.SetData("def")
	c2.SetDone(false)

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
	u := unit.NewTodo()
	c1 := u.NewItem()
	c1.SetData("abc")
	c1.SetDone(true)
	c2 := u.NewItem()
	c2.SetData("def")
	c2.SetDone(false)
	u.SetItems([]unit.TodoItem{c1, c2})

	items := u.Items()

	tmp := items[0]
	assert.Equal(t, "abc", tmp.Data())
	assert.Equal(t, true, tmp.Done())

	tmp = items[1]
	assert.Equal(t, "def", tmp.Data())
	assert.Equal(t, false, tmp.Done())
}

func TestTodoItem_MarshalJSON(t *testing.T) {
	i := unit.NewTodo().NewItem()
	i.SetData("abc")
	i.SetDone(true)

	bytes, err := json.Marshal(i)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"data": "abc", "done": true}`, string(bytes))
}

func TestTodoItem_UnmarshalJSON(t *testing.T) {
	i := unit.NewTodo().NewItem()

	err := json.Unmarshal([]byte(`{"data": "abc", "done": true}`), &i)
	assert.NoError(t, err)
	assert.Equal(t, "abc", i.Data())
	assert.True(t, i.Done())
}

func TestTodo_MarshalJSON(t *testing.T) {
	u := unit.NewTodo(
		unit.OptionClockMock(createdTime, updatedTime),
		unit.OptionTitle("MyUnit"),
	)
	c1 := u.NewItem()
	c1.SetData("abc")
	c1.SetDone(true)
	c2 := u.NewItem()
	c2.SetData("def")
	c2.SetDone(false)
	u.SetItems([]unit.TodoItem{c1, c2})

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, toJSON(jsonUnit(u), `"items":[{"data": "abc", "done": true},{"data": "def", "done": false}]`), string(bytes))
}

func TestTodo_UnmarshalJSON(t *testing.T) {
	u := unit.NewTodo()

	err := json.Unmarshal([]byte(toJSON(jsonUnit(u), `"items":[{"data": "abc", "done": true},{"data": "def", "done": false}]`)), &u)
	assert.NoError(t, err)

	items := u.Items()
	assert.Len(t, items, 2)

	tmp := items[0]
	assert.Equal(t, "abc", tmp.Data())
	assert.Equal(t, true, tmp.Done())

	tmp = items[1]
	assert.Equal(t, "def", tmp.Data())
	assert.Equal(t, false, tmp.Done())
}
