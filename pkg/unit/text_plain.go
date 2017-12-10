package unit

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
func (t *baseTextPlain) Data() string {
	return t.data
}

// SetData sets new unit data
func (t *baseTextPlain) SetData(data string) {
	t.data = data
}
