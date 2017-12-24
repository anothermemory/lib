package unit

import "encoding/json"

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

// NewList creates new List unit with given title
func NewList() List {
	return newBaseList()
}

func newBaseList() *baseList {
	return &baseList{baseUnit: *newBaseUnit(TypeList)}
}

// Items returns unit child units
func (u *baseList) Items() []Unit {
	return u.items
}

// SetItems sets unit child units
func (u *baseList) SetItems(items []Unit) {
	u.items = items
	u.refreshUpdated()
}

// AddItem adds new child unit to the unit
func (u *baseList) AddItem(item Unit) {
	u.items = append(u.items, item)
	u.refreshUpdated()
}

// GetItem returns child unit with given index
func (u *baseList) GetItem(index int) Unit {
	return u.items[index]
}

// SetItem sets child unit with given index to new unit
func (u *baseList) SetItem(index int, item Unit) {
	u.items[index] = item
	u.refreshUpdated()
}

// RemoveItem removes child unit with given index
func (u *baseList) RemoveItem(index int) {
	u.items = append(u.items[:index], u.items[index+1:]...)
	u.refreshUpdated()
}

func (u *baseList) Type() Type {
	return TypeList
}

type baseListJSON struct {
	baseUnitJSON
	Items []json.RawMessage `json:"items"`
}

type baseListItemJSON struct {
	Type Type `json:"type"`
}

func (u *baseList) fromJSONStruct(j baseListJSON) error {
	for _, i := range j.Items {
		item, err := createListItemFromJSON(i)
		if err != nil {
			continue
		}
		u.AddItem(item)
	}

	return nil
}

func createListItemFromJSON(j json.RawMessage) (Unit, error) {
	var ji baseListItemJSON
	err := json.Unmarshal(j, &ji)
	if err != nil {
		return nil, err
	}

	i := ji.Type.NewObject()

	err = json.Unmarshal(j, &i)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (u *baseList) MarshalJSON() ([]byte, error) {
	var items []json.RawMessage
	for _, i := range u.items {
		ij, err := json.Marshal(i)
		if err != nil {
			continue
		}
		items = append(items, json.RawMessage(ij))
	}
	return json.Marshal(baseListJSON{
		baseUnitJSON: baseUnitJSON{ID: u.id, Title: u.title, Type: u.unitType, Created: u.created, Updated: u.updated},
		Items:        items,
	})
}

func (u *baseList) UnmarshalJSON(b []byte) error {
	err := u.baseUnit.UnmarshalJSON(b)
	if err != nil {
		return err
	}

	var jsonData baseListJSON
	err = json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return u.fromJSONStruct(jsonData)
}
