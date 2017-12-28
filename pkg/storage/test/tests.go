package test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/storage"
	"github.com/stretchr/testify/assert"
)

type CreateFunc func() storage.Storage
type Func func(t *testing.T, c CreateFunc)

func RunStorageTests(t *testing.T, c CreateFunc) {
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) { test.testFunc(t, c) })
	}

}

var tests = []struct {
	title    string
	testFunc Func
}{
	{"Storage is not created initially when initialized first time with given arguments", func(t *testing.T, c CreateFunc) {
		s := c()
		assert.NotNil(t, s)
		assert.False(t, s.IsCreated())
	}},
}
