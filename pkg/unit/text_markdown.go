package unit

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
	return &baseTextMarkdown{baseTextPlain: baseTextPlain{data: data, baseUnit: baseUnit{title: title}}}
}

func (t *baseTextMarkdown) Render() string {
	panic("implement me")
}
