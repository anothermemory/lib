package unit

import "encoding/json"

// TextCode represents unit which contains some source code as plain text and it's language
type TextCode interface {
	TextPlain
	Language() string
	SetLanguage(language string)
}

// baseTextCode represents default implementation of TextCode interface
type baseTextCode struct {
	baseTextPlain
	language string
}

// NewTextCode creates new TextCode unit with given title, data and language
func NewTextCode(options ...func(u interface{})) TextCode {
	u := &baseTextCode{baseTextPlain: baseTextPlain{baseUnit: *newBaseUnit(TypeTextCode)}}

	initUnit(&u.baseUnit, options...)
	initUnit(&u.baseTextPlain, options...)
	initUnit(u, options...)

	return u
}

// Language returns unit language
func (u *baseTextCode) Language() string {
	return u.language
}

// SetLanguage sets new unit language
func (u *baseTextCode) SetLanguage(language string) {
	u.language = language
	u.refreshUpdated()
}

type baseTextCodeJSON struct {
	baseTextPlainJSON
	Language string `json:"language"`
}

func (u *baseTextCode) fromJSONStruct(j baseTextCodeJSON) error {
	u.language = j.Language

	return nil
}

func (u *baseTextCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(baseTextCodeJSON{
		baseTextPlainJSON: baseTextPlainJSON{
			baseUnitJSON: baseUnitJSON{ID: u.id, Title: u.title, Type: u.unitType, Created: u.created, Updated: u.updated},
			Data:         u.Data(),
		}, Language: u.language})
}

func (u *baseTextCode) UnmarshalJSON(b []byte) error {
	err := u.baseTextPlain.UnmarshalJSON(b)
	if err != nil {
		return err
	}

	var jsonData baseTextCodeJSON
	err = json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return u.fromJSONStruct(jsonData)
}

// TextCodeData is an option that sets data for a text code unit to the provided value
func TextCodeData(t string) func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseTextCode); converted {
			o.data = t
		}
	}
}

// TextCodeLanguage is an option that sets language for a text code unit to the provided value
func TextCodeLanguage(l string) func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseTextCode); converted {
			o.language = l
		}
	}
}
