package unit_test

import (
	"testing"

	"encoding/json"

	"fmt"

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
	assert.NotEmpty(t, u.ID())
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

func TestTodoItem_MarshalJSON(t *testing.T) {
	i := unit.NewTodoItem("abc", true)

	bytes, err := json.Marshal(i)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"data": "abc", "done": true}`, string(bytes))
}

func TestTodoItem_UnmarshalJSON(t *testing.T) {
	i := unit.NewTodoItem("", false)

	err := json.Unmarshal([]byte(`{"data": "abc", "done": true}`), &i)
	assert.NoError(t, err)
	assert.Equal(t, "abc", i.Data())
	assert.True(t, i.Done())
}

func TestTodo_Type(t *testing.T) {
	assert.Equal(t, unit.TypeTodo, unit.NewTodo("").Type())
}

func TestTodo_MarshalJSON(t *testing.T) {
	u := unit.NewTodo("MyUnit")
	u.SetItems([]unit.TodoItem{
		unit.NewTodoItem("abc", true),
		unit.NewTodoItem("def", false)})

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, fmt.Sprintf(`{"id": "%s", "title": "MyUnit", "type":"todo", "items":[{"data": "abc", "done": true},{"data": "def", "done": false}]}`, u.ID()), string(bytes))
}

func TestTodo_UnmarshalJSON(t *testing.T) {
	u := unit.NewTodo("MyUnit")

	err := json.Unmarshal([]byte(`{"id": "123", "title": "MyUnit", "type":"todo", "items":[{"data": "abc", "done": true},{"data": "def", "done": false}]}`), &u)
	assert.NoError(t, err)
	assert.Equal(t, "123", u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, unit.TypeTodo, u.Type())

	items := u.Items()
	assert.Len(t, items, 2)

	tmp := items[0]
	assert.Equal(t, "abc", tmp.Data())
	assert.Equal(t, true, tmp.Done())

	tmp = items[1]
	assert.Equal(t, "def", tmp.Data())
	assert.Equal(t, false, tmp.Done())
}
