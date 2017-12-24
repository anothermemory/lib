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
func NewTextCode() TextCode {
	return newBaseTextCode()
}

func newBaseTextCode() *baseTextCode {
	return &baseTextCode{baseTextPlain: baseTextPlain{baseUnit: *newBaseUnit(TypeTextCode)}}
}

// Language returns unit language
func (u *baseTextCode) Language() string {
	return u.language
}

// SetLanguage sets new unit language
func (u *baseTextCode) SetLanguage(language string) {
	u.language = language
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
