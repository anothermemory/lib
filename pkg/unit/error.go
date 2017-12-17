package unit

import "fmt"

type JSONTypeError struct {
	Expected, Actual string
}

func (e JSONTypeError) Error() string {
	return fmt.Sprintf("Unexpected unit type in json data received. Expected: %s, Actual: %s", e.Expected, e.Actual)
}

type TypeError struct {
	Type string
}

func (e TypeError) Error() string {
	return fmt.Sprintf("Unexpected type: %s", e.Type)
}
