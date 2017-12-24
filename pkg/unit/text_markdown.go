package unit

// TextMarkdown represents unit which has some plain markdown content
type TextMarkdown interface {
	TextPlain
}

// baseTextMarkdown represents default implementation of TextMarkdown interface
type baseTextMarkdown struct {
	baseTextPlain
}

// NewTextMarkdown creates new TextMarkdown unit with given title and data
func NewTextMarkdown() TextMarkdown {
	return newBaseTextMarkdown()
}

func newBaseTextMarkdown() *baseTextMarkdown {
	return &baseTextMarkdown{baseTextPlain: baseTextPlain{baseUnit: *newBaseUnit(TypeTextMarkdown)}}
}
