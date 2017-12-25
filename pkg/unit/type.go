package unit

import (
	"encoding/json"
	"fmt"
)

// Type represents type enumeration for unit types
type Type int

const (
	// TypeUnit represents type "Unit" for unit
	TypeUnit Type = 1 + iota
	// TypeList represents type "List" for unit
	TypeList
	// TypeTodo represents type "Todo" for unit
	TypeTodo
	// TypeTextPlain represents type "TextPlain" for unit
	TypeTextPlain
	// TypeTextMarkdown represents type "TextMarkdown" for unit
	TypeTextMarkdown
	// TypeTextCode represents type "TextCode" for unit
	TypeTextCode
)

var typeStrings = map[Type]string{
	TypeUnit:         "unit",
	TypeList:         "list",
	TypeTodo:         "todo",
	TypeTextPlain:    "text_plain",
	TypeTextMarkdown: "text_markdown",
	TypeTextCode:     "text_code",
}

var typeObjects = map[Type]func(options ...func(u interface{})) Unit{
	TypeUnit:         func(options ...func(u interface{})) Unit { return NewUnit(options...) },
	TypeList:         func(options ...func(u interface{})) Unit { return NewList(options...) },
	TypeTodo:         func(options ...func(u interface{})) Unit { return NewTodo(options...) },
	TypeTextPlain:    func(options ...func(u interface{})) Unit { return NewTextPlain(options...) },
	TypeTextMarkdown: func(options ...func(u interface{})) Unit { return NewTextMarkdown(options...) },
	TypeTextCode:     func(options ...func(u interface{})) Unit { return NewTextCode(options...) },
}

// String returns string representation of type
func (t Type) String() string {
	val, ok := typeStrings[t]
	if !ok {
		return fmt.Sprintf("Type(%d)", t)
	}
	return val
}

// TypeFromString returns type by it's string representation
func TypeFromString(s string) (Type, error) {
	for k, v := range typeStrings {
		if v == s {
			return k, nil
		}
	}

	return 0, fmt.Errorf("%s does not belong to Type values", s)
}

// MarshalJSON converts type to it's json representation (normally string)
func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON converts json type representation (normally string) to type object
func (t *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("type should be a string, got %s", data)
	}

	var err error
	*t, err = TypeFromString(s)
	return err
}

// NewObject creates new empty object of given type
func (t *Type) NewObject(options ...func(u interface{})) Unit {
	return typeObjects[*t](options...)
}
