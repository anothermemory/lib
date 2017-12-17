package unit

import (
	"encoding/json"

	"github.com/russross/blackfriday"
)

// TextMarkdown represents unit which has some plain markdown content
type TextMarkdown interface {
	TextPlain
	Render() string
}

// baseTextMarkdown represents default implementation of TextMarkdown interface
type baseTextMarkdown struct {
	baseTextPlain
}

// NewTextMarkdown creates new TextMarkdown unit with given title and data
func NewTextMarkdown(title string, data string) TextMarkdown {
	return &baseTextMarkdown{baseTextPlain: baseTextPlain{data: data, baseUnit: *newBaseUnit(title)}}
}

func (u *baseTextMarkdown) Render() string {
	return string(blackfriday.Run([]byte(u.data)))
}

func (u *baseTextMarkdown) Type() string {
	return "text_markdown"
}

type baseTextMarkdownJSON struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Data  string `json:"data"`
}

func (u *baseTextMarkdown) fromJSONStruct(j baseTextMarkdownJSON) error {
	if j.Type != u.Type() {
		return JSONTypeError{Expected: u.Type(), Actual: j.Type}
	}
	u.SetID(j.ID)
	u.SetTitle(j.Title)
	u.SetData(j.Data)

	return nil
}

func (u *baseTextMarkdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(baseTextMarkdownJSON{ID: u.ID(), Title: u.Title(), Type: u.Type(), Data: u.Data()})
}

func (u *baseTextMarkdown) UnmarshalJSON(b []byte) error {
	var jsonData baseTextMarkdownJSON
	err := json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return u.fromJSONStruct(jsonData)
}
