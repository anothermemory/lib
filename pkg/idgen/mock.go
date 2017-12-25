package idgen

type generatorMock struct {
	value string
}

func (g generatorMock) Generate() string {
	return g.value
}

// NewMock returns Generator which will always return passed value
func NewMock(value string) Generator {
	return &generatorMock{value: value}
}
