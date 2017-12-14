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
}

// baseUnit represents default implementation of Unit interface
type baseUnit struct {
	Unit
	id    string
	title string
}

// NewUnit creates new Unit with given title. Unit id is generated automatically
func NewUnit(title string) Unit {
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

func (u *baseUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}{
		ID:    u.ID(),
		Title: u.Title(),
	})
}

func (u *baseUnit) UnmarshalJSON(b []byte) error {
	var jsonData = struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}{}

	err := json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	u.SetID(jsonData.ID)
	u.SetTitle(jsonData.Title)

	return nil
}
