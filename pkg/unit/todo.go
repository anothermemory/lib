package unit

import (
	"encoding/json"
)

// Todo represents unit which contains todo items
type Todo interface {
	Unit
	NewItem() TodoItem
	Items() []TodoItem
	SetItems(items []TodoItem)
	AddItem(item TodoItem)
	GetItem(index int) TodoItem
	SetItem(index int, item TodoItem)
	RemoveItem(index int)
}

// TodoItem represents single item which can have some string data and done status
type TodoItem interface {
	Data() string
	SetData(data string)
	Done() bool
	SetDone(done bool)
}

// baseTodo represents default implementation of Todo interface
type baseTodo struct {
	baseUnit
	items []TodoItem
}

// baseTodo represents default implementation of TodoItem interface
type baseTodoItem struct {
	todo *baseTodo
	data string
	done bool
}

// Data returns item data
func (i *baseTodoItem) Data() string {
	return i.data
}

// SetData sets new item
func (i *baseTodoItem) SetData(data string) {
	i.data = data
	i.todo.refreshUpdated()
}

// Done returns item done status
func (i *baseTodoItem) Done() bool {
	return i.done
}

// SetDone sets new item done status
func (i *baseTodoItem) SetDone(done bool) {
	i.done = done
	i.todo.refreshUpdated()
}

type baseTodoItemJSON struct {
	Data string `json:"data"`
	Done bool   `json:"done"`
}

func (i *baseTodoItem) fromJSONStruct(j baseTodoItemJSON) error {
	i.data = j.Data
	i.done = j.Done

	return nil
}

func (i *baseTodoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(baseTodoItemJSON{Data: i.Data(), Done: i.Done()})
}

func (i *baseTodoItem) UnmarshalJSON(b []byte) error {
	var jsonData baseTodoItemJSON
	err := json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return i.fromJSONStruct(jsonData)
}

// NewTodo creates new Todo unit with given title
func NewTodo(options ...func(u interface{})) Todo {
	u := &baseTodo{baseUnit: *newBaseUnit(TypeTodo)}

	initUnitOptions(&u.baseUnit, options...)
	initUnitOptions(u, options...)
	initUnit(&u.baseUnit)

	return u
}

func (u *baseTodo) NewItem() TodoItem {
	return &baseTodoItem{todo: u}
}

// Items returns unit child items
func (u *baseTodo) Items() []TodoItem {
	return u.items
}

// SetItems sets unit child items
func (u *baseTodo) SetItems(items []TodoItem) {
	u.items = items
	u.refreshUpdated()
}

// AddItem adds new child item to the unit
func (u *baseTodo) AddItem(item TodoItem) {
	u.items = append(u.items, item)
	u.refreshUpdated()
}

// GetItem returns child item with given index
func (u *baseTodo) GetItem(index int) TodoItem {
	return u.items[index]
}

// SetItem sets child item with given index to new item
func (u *baseTodo) SetItem(index int, item TodoItem) {
	u.items[index] = item
	u.refreshUpdated()
}

// RemoveItem removes child item with given index
func (u *baseTodo) RemoveItem(index int) {
	u.items = append(u.items[:index], u.items[index+1:]...)
	u.refreshUpdated()
}

type baseTodoJSON struct {
	baseUnitJSON
	Items []baseTodoItemJSON `json:"items"`
}

func (u *baseTodo) fromJSONStruct(j baseTodoJSON) error {
	for _, v := range j.Items {
		u.AddItem(&baseTodoItem{todo: u, data: v.Data, done: v.Done})
	}

	return nil
}

func (u *baseTodo) MarshalJSON() ([]byte, error) {
	var items []baseTodoItemJSON
	for _, v := range u.items {
		items = append(items, baseTodoItemJSON{Data: v.Data(), Done: v.Done()})
	}
	return json.Marshal(baseTodoJSON{
		baseUnitJSON: baseUnitJSON{ID: u.id, Title: u.title, Type: u.unitType, Created: u.created, Updated: u.updated},
		Items:        items,
	})
}

func (u *baseTodo) UnmarshalJSON(b []byte) error {
	err := u.baseUnit.UnmarshalJSON(b)
	if err != nil {
		return err
	}

	var jsonData baseTodoJSON
	err = json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return u.fromJSONStruct(jsonData)
}

// OptionTodoItem is an option that adds new todo item to the unit
func OptionTodoItem(data string, done bool) func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseTodo); converted {
			o.items = append(o.items, &baseTodoItem{todo: o, data: data, done: done})
		}
	}
}
