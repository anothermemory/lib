package storage

import (
	"errors"
	"fmt"
	"strings"
)

// UnitNotFoundError represents error where some unit cannot be loaded by storage
type UnitNotFoundError struct {
	id string
}

func (u UnitNotFoundError) Error() string {
	return fmt.Sprintf("Unit with id (%s) was not found in storage and cannot be loaded", u.id)
}

// UnitsNotFoundError represents error where multiple units cannot be loaded by storage
type UnitsNotFoundError struct {
	ids []string
}

func (u UnitsNotFoundError) Error() string {
	return fmt.Sprintf("Units with following ids was not found in storage and cannot be loaded: (%s)", strings.Join(u.ids, ","))
}

// ErrNotCreated represents error where storage is tried to be used before it was created
var ErrNotCreated = errors.New("storage is not created yet and cannot be used")

//type JsonConfigError struct {
//	details string
//}
//
//func (j JsonConfigError) Error() string {
//	return fmt.Sprintf("Failed to parse storage json config: %s", j.details)
//}

// UnsupportedUnitError represents error where some unsupported unit is tried to be processed by storage
type UnsupportedUnitError struct {
	details string
}

func (u UnsupportedUnitError) Error() string {
	return fmt.Sprintf("Unit serialization is not supported for unit: %s", u.details)
}
