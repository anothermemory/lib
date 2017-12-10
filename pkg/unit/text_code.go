package unit

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
func (t *baseTextCode) Language() string {
	return t.language
}

// SetLanguage sets new unit language
func (t *baseTextCode) SetLanguage(language string) {
	t.language = language
}
