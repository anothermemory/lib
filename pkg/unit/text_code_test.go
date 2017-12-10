package unit_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextCode(t *testing.T) {
	u := unit.NewTextCode("MyUnit", "MyData", "PHP")
	assert.NotNil(t, u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, "MyData", u.Data())
	assert.Equal(t, "PHP", u.Language())
}

func TestTextCode_Language(t *testing.T) {
	u := unit.NewTextCode("MyUnit", "MyData", "PHP")
	u.SetLanguage("Java")
	assert.Equal(t, "Java", u.Language())
}
