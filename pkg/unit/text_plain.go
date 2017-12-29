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
func NewTextPlain(options ...func(u interface{})) TextPlain {
	u := &baseTextPlain{baseUnit: *newBaseUnit(TypeTextPlain)}
	initUnitOptions(&u.baseUnit, options...)
	initUnitOptions(u, options...)
	initUnit(&u.baseUnit)

	return u
}

// Data returns unit data
func (u *baseTextPlain) Data() string {
	return u.data
}

// SetData sets new unit data
func (u *baseTextPlain) SetData(data string) {
	u.data = data
	u.refreshUpdated()
}

type baseTextPlainJSON struct {
	baseUnitJSON
	Data string `json:"data"`
}

func (u *baseTextPlain) fromJSONStruct(j baseTextPlainJSON) error {
	u.data = j.Data
	return nil
}

func (u *baseTextPlain) MarshalJSON() ([]byte, error) {
	return json.Marshal(baseTextPlainJSON{
		baseUnitJSON: baseUnitJSON{ID: u.id, Title: u.title, Type: u.unitType, Created: u.created, Updated: u.updated},
		Data:         u.Data(),
	})
}

func (u *baseTextPlain) UnmarshalJSON(b []byte) error {
	err := u.baseUnit.UnmarshalJSON(b)
	if err != nil {
		return err
	}

	var jsonData baseTextPlainJSON
	err = json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return u.fromJSONStruct(jsonData)
}

// OptionTextPlainData is an option that sets data for a text plain unit to the provided value
func OptionTextPlainData(t string) func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseTextPlain); converted {
			o.data = t
		}
	}
}
