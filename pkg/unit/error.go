package unit

import "fmt"

// JSONTypeError represents error when type received in JSON message is different from type of the target unit
type JSONTypeError struct {
	Expected, Actual string
}

func (e JSONTypeError) Error() string {
	return fmt.Sprintf("Unexpected unit type in json data received. Expected: %s, Actual: %s", e.Expected, e.Actual)
}

// TypeError represents error when type of some unit is unexpected at this point
type TypeError struct {
	Type string
}

func (e TypeError) Error() string {
	return fmt.Sprintf("Unexpected type: %s", e.Type)
}
