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
func (t *baseTodo) Items() []TodoItem {
	return t.items
}

// SetItems sets unit child items
func (t *baseTodo) SetItems(items []TodoItem) {
	t.items = items
}

// AddItem adds new child item to the unit
func (t *baseTodo) AddItem(item TodoItem) {
	t.items = append(t.items, item)
}

// GetItem returns child item with given index
func (t *baseTodo) GetItem(index int) TodoItem {
	return t.items[index]
}

// SetItem sets child item with given index to new item
func (t *baseTodo) SetItem(index int, item TodoItem) {
	t.items[index] = item
}

// RemoveItem removes child item with given index
func (t *baseTodo) RemoveItem(index int) {
	t.items = append(t.items[:index], t.items[index+1:]...)
}

// NewTodoItem creates new TodoItem with given data and status
func NewTodoItem(data string, done bool) TodoItem {
	return &baseTodoItem{data: data, done: done}
}

// Data returns item data
func (t *baseTodoItem) Data() string {
	return t.data
}

// SetData sets new item
func (t *baseTodoItem) SetData(data string) {
	t.data = data
}

// Done returns item done status
func (t *baseTodoItem) Done() bool {
	return t.done
}

// SetDone sets new item done status
func (t *baseTodoItem) SetDone(done bool) {
	t.done = done
}
