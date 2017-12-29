package test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/storage"
	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/require"
)

// CreateFunc represents function which must return created storage object
type CreateFunc func() storage.Storage

// Func represents test function for single test-case
type Func func(t *testing.T, c CreateFunc, is *require.Assertions)

// RunStorageTests performs full test run for all test-cases for given storage
func RunStorageTests(t *testing.T, c CreateFunc) {
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) { test.testFunc(t, c, require.New(t)) })
	}

}

var tests = []struct {
	title    string
	testFunc Func
}{
	{"Storage is not created initially when initialized first time with given arguments", func(t *testing.T, c CreateFunc, is *require.Assertions) {
		is.False(c().IsCreated())
	}},
	{"Storage can be successfully created", func(t *testing.T, c CreateFunc, is *require.Assertions) {
		s := c()
		is.NoError(s.Create())
		is.True(s.IsCreated())
	}},
	{"Storage can not be used before it will be created", func(t *testing.T, c CreateFunc, is *require.Assertions) {
		s := c()
		u := unit.NewUnit()
		is.Error(s.SaveUnit(u))
		is.Error(s.RemoveUnit(u))
		u, e := s.LoadUnit("123")
		is.Error(e)
		is.Nil(u)
	}},
}
