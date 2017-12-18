package unit_test

import (
	"testing"

	"io/ioutil"
	"path/filepath"

	"encoding/json"
	"fmt"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestNewTextMarkdown(t *testing.T) {
	u := unit.NewTextMarkdown("MyUnit", "MyData")
	assert.NotNil(t, u.ID())
	assert.NotEmpty(t, u.ID())
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

func TestTextMarkdown_MarshalJSON(t *testing.T) {
	u := unit.NewTextMarkdown("MyUnit", "abc")

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, fmt.Sprintf(`{"id": "%s", "title": "MyUnit", "data": "abc", "type": "text_markdown"}`, u.ID()), string(bytes))
}

func TestTextMarkdown_UnmarshalJSON(t *testing.T) {
	u := unit.NewTextMarkdown("", "")

	err := json.Unmarshal([]byte(`{"id": "123", "title": "MyUnit", "data": "abc", "type": "text_markdown"}`), &u)
	assert.NoError(t, err)
	assert.Equal(t, "123", u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, "abc", u.Data())
	assert.Equal(t, unit.TypeTextMarkdown, u.Type())
}

func TestTextMarkdown_Type(t *testing.T) {
	assert.Equal(t, unit.TypeTextMarkdown, unit.NewTextMarkdown("", "").Type())
}

func readFile(t *testing.T, name string) string {
	filename := filepath.Join("testdata", name)
	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
	}
	return string(inputBytes)
}
