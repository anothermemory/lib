package unit

// Todo represents unit which contains todo items
type Todo interface {
	Unit
	Items() []TodoItem
	SetItems(items []TodoItem)
	AddItem(item TodoItem)
	GetItem(index int) TodoItem
	SetItem(index int, item TodoItem)
	RemoveItem(index int)
}

// baseTodo represents default implementation of Todo interface
type baseTodo struct {
	baseUnit
	items []TodoItem
}

// TodoItem represents single item which can have some string data and done status
type TodoItem interface {
	Data() string
	SetData(data string)
	Done() bool
	SetDone(done bool)
}

// baseTodo represents default implementation of TodoItem interface
type baseTodoItem struct {
	TodoItem
	data string
	done bool
}

// NewTodo creates new Todo unit with given title
func NewTodo(title string) Todo {
	return &baseTodo{baseUnit: baseUnit{title: title}}
}

// Items returns unit child items
func (u *baseTodo) Items() []TodoItem {
	return u.items
}

// SetItems sets unit child items
func (u *baseTodo) SetItems(items []TodoItem) {
	u.items = items
}

// AddItem adds new child item to the unit
func (u *baseTodo) AddItem(item TodoItem) {
	u.items = append(u.items, item)
}

// GetItem returns child item with given index
func (u *baseTodo) GetItem(index int) TodoItem {
	return u.items[index]
}

// SetItem sets child item with given index to new item
func (u *baseTodo) SetItem(index int, item TodoItem) {
	u.items[index] = item
}

// RemoveItem removes child item with given index
func (u *baseTodo) RemoveItem(index int) {
	u.items = append(u.items[:index], u.items[index+1:]...)
}

func (u *baseTodo) Type() string {
	return "todo"
}

// NewTodoItem creates new TodoItem with given data and status
func NewTodoItem(data string, done bool) TodoItem {
	return &baseTodoItem{data: data, done: done}
}

// Data returns item data
func (i *baseTodoItem) Data() string {
	return i.data
}

// SetData sets new item
func (i *baseTodoItem) SetData(data string) {
	i.data = data
}

// Done returns item done status
func (i *baseTodoItem) Done() bool {
	return i.done
}

// SetDone sets new item done status
func (i *baseTodoItem) SetDone(done bool) {
	i.done = done
}
