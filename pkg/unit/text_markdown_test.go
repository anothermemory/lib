package unit_test

import (
	"testing"

	"io/ioutil"
	"path/filepath"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextMarkdown(t *testing.T) {
	u := unit.NewTextMarkdown("MyUnit", "MyData")
	assert.NotNil(t, u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, "MyData", u.Data())
}

func TestTextMarkdown_Render(t *testing.T) {
	cases := []string{
		"Markdown Documentation - Basics",
	}

	for _, tc := range cases {
		input := readFile(t, tc+".text")
		expected := readFile(t, tc+".html")

		actual := unit.NewTextMarkdown("MyUnit", input).Render()
		assert.Equal(t, expected, actual, tc)
	}
}

func readFile(t *testing.T, name string) string {
	filename := filepath.Join("testdata", name)
	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
	}
	return string(inputBytes)
}
