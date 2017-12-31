package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestType_String(t *testing.T) {
	tests := []struct {
		unitType       unit.Type
		expectedString string
	}{
		{unit.TypeUnit, "unit"},
		{unit.TypeTextPlain, "text_plain"},
		{unit.TypeTextMarkdown, "text_markdown"},
		{unit.TypeTextCode, "text_code"},
		{unit.TypeTodo, "todo"},
		{unit.TypeList, "list"},
		{unit.Type(123), "Type(123)"},
	}
	for _, test := range tests {
		t.Run(test.unitType.String(), func(t *testing.T) {
			assert.Equal(t, test.expectedString, test.unitType.String())
		})
	}
}

func TestType_FromString(t *testing.T) {
	tests := []struct {
		expectedType unit.Type
		inputString  string
	}{
		{unit.TypeUnit, "unit"},
		{unit.TypeTextPlain, "text_plain"},
		{unit.TypeTextMarkdown, "text_markdown"},
		{unit.TypeTextCode, "text_code"},
		{unit.TypeTodo, "todo"},
		{unit.TypeList, "list"},
	}
	for _, test := range tests {
		t.Run(test.expectedType.String(), func(t *testing.T) {
			resultType, err := unit.TypeFromString(test.inputString)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedType, resultType)
		})
	}
}

func TestType_FromString_Unexpected(t *testing.T) {
	resultType, err := unit.TypeFromString("1234")
	assert.Error(t, err)
	assert.NotNil(t, resultType)
}

func TestType_UnmarshalJSON_Malformed(t *testing.T) {
	unitType := unit.TypeUnit
	err := json.Unmarshal([]byte("123"), &unitType)
	assert.Error(t, err)
}

func TestType_NewObject(t *testing.T) {
	tests := []unit.Type{
		unit.TypeUnit,
		unit.TypeTextPlain,
		unit.TypeTextMarkdown,
		unit.TypeTextCode,
		unit.TypeTodo,
		unit.TypeList,
	}
	for _, test := range tests {
		t.Run(test.String(), func(t *testing.T) {
			resultObject := test.NewObject()
			assert.NotNil(t, resultObject)
			assert.Equal(t, test, resultObject.Type())
		})
	}
}
