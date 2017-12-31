package unit_test

import (
	"testing"
	"time"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	mockDateTime := time.Date(2017, 12, 30, 17, 52, 00, 0, time.Local)
	mockAnotherTime := time.Date(2016, 12, 30, 17, 52, 00, 0, time.Local)
	updatedUnit := unit.NewUnit(
		unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime, mockAnotherTime),
	)
	updatedUnit.SetTitle("Title")

	tests := []struct {
		name   string
		u1, u2 unit.Unit
		result bool
	}{
		// Nil
		{"If unit 1 is nil then units are not equal", nil, unit.NewUnit(), false},
		{"If unit 2 is nil then units are not equal", unit.NewUnit(), nil, false},

		// Units
		{"New units are not equal", unit.NewUnit(), unit.NewUnit(), false},
		{"Units with the same ID,Type,Title,Created,Updated are equal", unit.NewUnit(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
		), unit.NewUnit(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
		), true},
		{"Units with different id are not equal", unit.NewUnit(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
		), unit.NewUnit(
			unit.OptionIDGeneratorMock("456"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
		), false},
		{"Units with different title are not equal", unit.NewUnit(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
		), unit.NewUnit(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title2"), unit.OptionClockMock(mockDateTime),
		), false},
		{"Units with different type are not equal", unit.NewUnit(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
		), unit.NewTextPlain(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
		), false},
		{"Units with different created are not equal", unit.NewUnit(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
		), unit.NewUnit(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockAnotherTime),
		), false},
		{"Units with different updated are not equal", unit.NewUnit(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
		), updatedUnit, false},

		// Todo units
		{"New todo units are not equal", unit.NewTodo(), unit.NewTodo(), false},
		{"Todo units with same items in same order are equal", unit.NewTodo(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTodoItem("D1", true), unit.OptionTodoItem("D2", false),
		), unit.NewTodo(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTodoItem("D1", true), unit.OptionTodoItem("D2", false),
		), true},
		{"Todo units with different items are not equal", unit.NewTodo(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTodoItem("D1", true), unit.OptionTodoItem("D2", false),
		), unit.NewTodo(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTodoItem("D3", true), unit.OptionTodoItem("D3", false),
		), false},
		{"Todo units with different items count are not equal", unit.NewTodo(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTodoItem("D1", true), unit.OptionTodoItem("D2", false),
		), unit.NewTodo(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTodoItem("D1", true),
		), false},
		{"Todo units with same items in different order are not equal", unit.NewTodo(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTodoItem("D1", true), unit.OptionTodoItem("D2", false),
		), unit.NewTodo(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTodoItem("D2", false), unit.OptionTodoItem("D1", true),
		), false},

		// Text plain units
		{"New text plain units are not equal", unit.NewTextPlain(), unit.NewTextPlain(), false},
		{"Text plain units with the same data are equal", unit.NewTextPlain(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextPlainData("Data"),
		), unit.NewTextPlain(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextPlainData("Data"),
		), true},
		{"Text plain units with different data are not equal", unit.NewTextPlain(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextPlainData("Data1"),
		), unit.NewTextPlain(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextPlainData("Data2"),
		), false},

		// Text markdown units
		{"New text markdown units are not equal", unit.NewTextMarkdown(), unit.NewTextMarkdown(), false},
		{"Text markdown units with the same data are equal", unit.NewTextMarkdown(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextMarkdownData("Data"),
		), unit.NewTextMarkdown(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextMarkdownData("Data"),
		), true},
		{"Text markdown units with different data are not equal", unit.NewTextMarkdown(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextMarkdownData("Data1"),
		), unit.NewTextMarkdown(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextMarkdownData("Data2"),
		), false},

		// Text code units
		{"New text code units are not equal", unit.NewTextCode(), unit.NewTextCode(), false},
		{"Text code units with the same data and language are equal", unit.NewTextCode(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextCodeData("Data"), unit.OptionTextCodeLanguage("Lang"),
		), unit.NewTextCode(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextCodeData("Data"), unit.OptionTextCodeLanguage("Lang"),
		), true},
		{"Text code units with different data are not equal", unit.NewTextCode(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextCodeData("Data1"), unit.OptionTextCodeLanguage("Lang"),
		), unit.NewTextCode(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextCodeData("Data2"), unit.OptionTextCodeLanguage("Lang"),
		), false},
		{"Text code units with different language are not equal", unit.NewTextCode(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextCodeData("Data"), unit.OptionTextCodeLanguage("Lang1"),
		), unit.NewTextCode(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionTextCodeData("Data"), unit.OptionTextCodeLanguage("Lang2"),
		), false},

		// List units
		{"New list units are not equal", unit.NewList(), unit.NewList(), false},
		{"List units with same items in same order are equal", unit.NewList(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title1"), unit.OptionClockMock(mockDateTime))),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("456"), unit.OptionTitle("Title2"), unit.OptionClockMock(mockDateTime))),
		), unit.NewList(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title1"), unit.OptionClockMock(mockDateTime))),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("456"), unit.OptionTitle("Title2"), unit.OptionClockMock(mockDateTime))),
		), true},
		{"List units with different items are not equal", unit.NewList(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title1"), unit.OptionClockMock(mockDateTime))),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("456"), unit.OptionTitle("Title2"), unit.OptionClockMock(mockDateTime))),
		), unit.NewList(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("789"), unit.OptionTitle("Title1"), unit.OptionClockMock(mockDateTime))),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("789"), unit.OptionTitle("Title2"), unit.OptionClockMock(mockDateTime))),
		), false},
		{"List units with different items count are not equal", unit.NewList(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title1"), unit.OptionClockMock(mockDateTime))),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("456"), unit.OptionTitle("Title2"), unit.OptionClockMock(mockDateTime))),
		), unit.NewList(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("789"), unit.OptionTitle("Title1"), unit.OptionClockMock(mockDateTime))),
		), false},
		{"List units with same items in different order are not equal", unit.NewList(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title1"), unit.OptionClockMock(mockDateTime))),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("456"), unit.OptionTitle("Title2"), unit.OptionClockMock(mockDateTime))),
		), unit.NewList(
			unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title"), unit.OptionClockMock(mockDateTime),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("456"), unit.OptionTitle("Title2"), unit.OptionClockMock(mockDateTime))),
			unit.OptionListItem(unit.NewUnit(unit.OptionIDGeneratorMock("123"), unit.OptionTitle("Title1"), unit.OptionClockMock(mockDateTime))),
		), false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, unit.Equal(test.u1, test.u2))
		})
	}
}
