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
	return &baseTextCode{baseTextPlain: baseTextPlain{data: data, baseUnit: baseUnit{title: title}}, language: language}
}

// Language returns unit language
func (u *baseTextCode) Language() string {
	return u.language
}

// SetLanguage sets new unit language
func (u *baseTextCode) SetLanguage(language string) {
	u.language = language
}

func (u *baseTextCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Data     string `json:"data"`
		Language string `json:"language"`
	}{
		ID:       u.ID(),
		Title:    u.Title(),
		Data:     u.Data(),
		Language: u.Language(),
	})
}

func (u *baseTextCode) UnmarshalJSON(b []byte) error {
	var jsonData = struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Data     string `json:"data"`
		Language string `json:"language"`
	}{}

	err := json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	u.SetID(jsonData.ID)
	u.SetTitle(jsonData.Title)
	u.SetData(jsonData.Data)
	u.SetLanguage(jsonData.Language)

	return nil
}
