package storage

import (
	"github.com/anothermemory/lib/pkg/unit"
)

// Storage represents some storage which can be used to store units
type Storage interface {
	SaveUnit(u unit.Unit) error
	RemoveUnit(u unit.Unit) error
	LoadUnit(id string) (unit.Unit, error)
	IsCreated() bool
	Create() error
	Remove() error
}
