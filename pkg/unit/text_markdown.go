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
func NewTextMarkdown(options ...func(u interface{})) TextMarkdown {
	u := &baseTextMarkdown{baseTextPlain: baseTextPlain{baseUnit: *newBaseUnit(TypeTextMarkdown)}}

	initUnitOptions(&u.baseUnit, options...)
	initUnitOptions(&u.baseTextPlain, options...)
	initUnitOptions(u, options...)
	initUnit(&u.baseUnit)

	return u
}

// TextMarkdownData is an option that sets data for a text markdown unit to the provided value
func TextMarkdownData(t string) func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseTextMarkdown); converted {
			o.data = t
		}
	}
}
