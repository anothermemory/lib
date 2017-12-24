package unit

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestTodo_MarshalJSON(t *testing.T) {
	u := newBaseTodo()
	u.SetItems([]TodoItem{
		NewTodoItem("abc", true),
		NewTodoItem("def", false)})
	freezeUnitTime(&u.baseUnit, createdTime, updatedTime)

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, toJSON(jsonUnit(u), `"items":[{"data": "abc", "done": true},{"data": "def", "done": false}]`), string(bytes))
}

func TestTodo_UnmarshalJSON(t *testing.T) {
	u := NewTodo()

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
