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
func NewList(title string) List {
	return &baseList{baseUnit: *newBaseUnit(title)}
}

// Items returns unit child units
func (u *baseList) Items() []Unit {
	return u.items
}

// SetItems sets unit child units
func (u *baseList) SetItems(items []Unit) {
	u.items = items
}

// AddItem adds new child unit to the unit
func (u *baseList) AddItem(item Unit) {
	u.items = append(u.items, item)
}

// GetItem returns child unit with given index
func (u *baseList) GetItem(index int) Unit {
	return u.items[index]
}

// SetItem sets child unit with given index to new unit
func (u *baseList) SetItem(index int, item Unit) {
	u.items[index] = item
}

// RemoveItem removes child unit with given index
func (u *baseList) RemoveItem(index int) {
	u.items = append(u.items[:index], u.items[index+1:]...)
}

func (u *baseList) Type() string {
	return "list"
}

type baseListJSON struct {
	ID    string            `json:"id"`
	Title string            `json:"title"`
	Type  string            `json:"type"`
	Items []json.RawMessage `json:"items"`
}

type baseListItemJSON struct {
	Type string `json:"type"`
}

func (u *baseList) fromJSONStruct(j baseListJSON) error {
	if j.Type != u.Type() {
		return JSONTypeError{Expected: u.Type(), Actual: j.Type}
	}
	u.SetID(j.ID)
	u.SetTitle(j.Title)

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
	var err error
	var i Unit
	err = json.Unmarshal(j, &ji)
	if err != nil {
		return nil, err
	}

	i, err = newUnitByType(ji.Type)
	if err != nil {
		return nil, err
	}

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
	return json.Marshal(baseListJSON{ID: u.ID(), Title: u.Title(), Type: u.Type(), Items: items})
}

func (u *baseList) UnmarshalJSON(b []byte) error {
	var jsonData baseListJSON
	err := json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return u.fromJSONStruct(jsonData)
}
