package unit

// List represents unit which contains other units
type List interface {
	Unit
	Items() []Unit
	SetItems(items []Unit)
	AddItem(item Unit)
	GetItem(index int) Unit
	SetItem(index int, item Unit)
	RemoveItem(index int)
}

// baseList represents default implementation for List interface
type baseList struct {
	baseUnit
	items []Unit
}

// NewList cretes new List unit with given title
func NewList(title string) List {
	return &baseList{baseUnit: baseUnit{title: title}}
}

// Items returns unit child units
func (l *baseList) Items() []Unit {
	return l.items
}

// SetItems sets unit child units
func (l *baseList) SetItems(items []Unit) {
	l.items = items
}

// AddItem adds new child unit to the unit
func (l *baseList) AddItem(item Unit) {
	l.items = append(l.items, item)
}

// GetItem returns child unit with given index
func (l *baseList) GetItem(index int) Unit {
	return l.items[index]
}

// SetItem sets child unit with given index to new unit
func (l *baseList) SetItem(index int, item Unit) {
	l.items[index] = item
}

// RemoveItem removes child unit with given index
func (l *baseList) RemoveItem(index int) {
	l.items = append(l.items[:index], l.items[index+1:]...)
}
