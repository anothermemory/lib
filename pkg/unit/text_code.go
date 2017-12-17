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
func NewTextCode(title string, data string, language string) TextCode {
	return &baseTextCode{baseTextPlain: baseTextPlain{data: data, baseUnit: *newBaseUnit(title)}, language: language}
}

// Language returns unit language
func (u *baseTextCode) Language() string {
	return u.language
}

// SetLanguage sets new unit language
func (u *baseTextCode) SetLanguage(language string) {
	u.language = language
}

func (u *baseTextCode) Type() string {
	return "text_code"
}

type baseTextCodeJSON struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Data     string `json:"data"`
	Language string `json:"language"`
}

func (u *baseTextCode) fromJSONStruct(j baseTextCodeJSON) error {
	if j.Type != u.Type() {
		return JSONTypeError{Expected: u.Type(), Actual: j.Type}
	}
	u.SetID(j.ID)
	u.SetTitle(j.Title)
	u.SetData(j.Data)
	u.SetLanguage(j.Language)

	return nil
}

func (u *baseTextCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(baseTextCodeJSON{ID: u.ID(), Title: u.Title(), Type: u.Type(), Data: u.Data(), Language: u.Language()})
}

func (u *baseTextCode) UnmarshalJSON(b []byte) error {
	var jsonData baseTextCodeJSON
	err := json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return u.fromJSONStruct(jsonData)
}
