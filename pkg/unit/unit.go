package unit

import (
	"encoding/json"

	"github.com/satori/go.uuid"
)

// Unit represents simplest unit which actually does nothing but used as a base for all other units
type Unit interface {
	json.Marshaler
	json.Unmarshaler
	ID() string
	SetID(id string)
	Title() string
	SetTitle(title string)
	Type() Type
}

// baseUnit represents default implementation of Unit interface
type baseUnit struct {
	Unit
	id    string
	title string
}

// NewUnit creates new Unit with given title. Unit id is generated automatically
func NewUnit(title string) Unit {
	return newBaseUnit(title)
}

func newBaseUnit(title string) *baseUnit {
	return &baseUnit{id: uuid.NewV4().String(), title: title}
}

// ID returns unit id
func (u *baseUnit) ID() string {
	return u.id
}

// SetTitle sets new unit title
func (u *baseUnit) SetID(id string) {
	u.id = id
}

// Title returns unit title
func (u *baseUnit) Title() string {
	return u.title
}

// SetTitle sets new unit title
func (u *baseUnit) SetTitle(title string) {
	u.title = title
}

func (u *baseUnit) Type() Type {
	return TypeUnit
}

type baseUnitJSON struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  Type   `json:"type"`
}

func (u *baseUnit) fromJSONStruct(j baseUnitJSON) error {
	if j.Type != u.Type() {
		return JSONTypeError{Expected: u.Type(), Actual: j.Type}
	}
	u.SetID(j.ID)
	u.SetTitle(j.Title)

	return nil
}

func (u *baseUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(baseUnitJSON{ID: u.ID(), Title: u.Title(), Type: u.Type()})
}

func (u *baseUnit) UnmarshalJSON(b []byte) error {
	var jsonData baseUnitJSON
	err := json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return u.fromJSONStruct(jsonData)
}

func newUnitByType(t Type) (Unit, error) {
	switch t {
	case TypeUnit:
		return NewUnit(""), nil
	case TypeTodo:
		return NewTodo(""), nil
	case TypeTextPlain:
		return NewTextPlain("", ""), nil
	case TypeTextMarkdown:
		return NewTextMarkdown("", ""), nil
	case TypeTextCode:
		return NewTextCode("", "", ""), nil
	case TypeList:
		return NewList(""), nil
	default:
		return nil, &TypeError{Type: t}

	}
}
