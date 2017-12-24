package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextCode(t *testing.T) {
	u := unit.NewTextCode()
	assert.Equal(t, unit.TypeTextCode, u.Type())
}

func TestBaseTextCode_Language(t *testing.T) {
	u := unit.NewTextCode()
	u.SetLanguage("Java")
	assert.Equal(t, "Java", u.Language())
}
