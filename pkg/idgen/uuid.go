package idgen

import "github.com/satori/go.uuid"

type generatorUUID struct {
}

func (generatorUUID) Generate() string {
	return uuid.NewV4().String()
}

// NewUUID returns Generator which will return new UUID V4 as generated values
func NewUUID() Generator {
	return &generatorUUID{}
}
