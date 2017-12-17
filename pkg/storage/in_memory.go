package storage

import "github.com/anothermemory/lib/pkg/unit"

type inMemoryStorage struct {
	units map[string]unit.Unit
}

// NewInMemoryStorage creates new storage which stores all units completely in memory
func NewInMemoryStorage() Storage {
	return &inMemoryStorage{}
}
func (s *inMemoryStorage) IsCreated() bool {
	return true
}

func (s *inMemoryStorage) Create() error {
	return nil
}

func (s *inMemoryStorage) Remove() error {
	return nil
}

func (s *inMemoryStorage) SaveUnit(u unit.Unit) error {
	switch u := u.(type) {
	case unit.List:
		s.units[u.ID()] = u
		for _, i := range u.Items() {
			s.units[i.ID()] = i
		}

	default:
		s.units[u.ID()] = u
	}
	return nil
}

func (s *inMemoryStorage) RemoveUnit(u unit.Unit) error {
	delete(s.units, u.ID())
	return nil
}

func (s *inMemoryStorage) LoadUnit(id string) (unit.Unit, error) {
	if u, ok := s.units[id]; ok {
		return u, nil
	}
	return nil, &UnitNotFoundError{id: id}
}

func (s *inMemoryStorage) LoadUnits(ids []string) (map[string]unit.Unit, error) {
	var result = make(map[string]unit.Unit)
	var errors []string
	for _, id := range ids {
		if u, ok := s.units[id]; ok {
			result[id] = u
		} else {
			errors = append(errors, id)
		}
	}

	if len(errors) > 0 {
		return result, UnitsNotFoundError{ids: errors}
	}

	return result, nil
}
