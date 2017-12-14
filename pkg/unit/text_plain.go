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
	return &baseTextPlain{baseUnit: baseUnit{title: title}, data: data}
}

// Data returns unit data
func (u *baseTextPlain) Data() string {
	return u.data
}

// SetData sets new unit data
func (u *baseTextPlain) SetData(data string) {
	u.data = data
}

func (u *baseTextPlain) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Data  string `json:"data"`
	}{
		ID:    u.ID(),
		Title: u.title,
		Data:  u.Data(),
	})
}

func (u *baseTextPlain) UnmarshalJSON(b []byte) error {
	var jsonData = struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Data  string `json:"data"`
	}{}

	err := json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	u.SetID(jsonData.ID)
	u.SetTitle(jsonData.Title)
	u.SetData(jsonData.Data)

	return nil
}
