package unit

import "encoding/json"

// TextPlain represents unit which has some plain text content
type TextPlain interface {
	Unit
	Data() string
	SetData(data string)
}

// baseTextPlain represents default implementation of TextPlain interface
type baseTextPlain struct {
	baseUnit
	data string
}

// NewTextPlain creates new TextPlain unit with given title and data
func NewTextPlain(title string, data string) TextPlain {
	return &baseTextPlain{baseUnit: *newBaseUnit(title), data: data}
}

// Data returns unit data
func (u *baseTextPlain) Data() string {
	return u.data
}

// SetData sets new unit data
func (u *baseTextPlain) SetData(data string) {
	u.data = data
}

func (u *baseTextPlain) Type() Type {
	return TypeTextPlain
}

type baseTextPlainJSON struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  Type   `json:"type"`
	Data  string `json:"data"`
}

func (u *baseTextPlain) fromJSONStruct(j baseTextPlainJSON) error {
	if j.Type != u.Type() {
		return JSONTypeError{Expected: u.Type(), Actual: j.Type}
	}
	u.SetID(j.ID)
	u.SetTitle(j.Title)
	u.SetData(j.Data)

	return nil
}

func (u *baseTextPlain) MarshalJSON() ([]byte, error) {
	return json.Marshal(baseTextPlainJSON{ID: u.ID(), Title: u.Title(), Type: u.Type(), Data: u.Data()})
}

func (u *baseTextPlain) UnmarshalJSON(b []byte) error {
	var jsonData baseTextPlainJSON
	err := json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return u.fromJSONStruct(jsonData)
}
