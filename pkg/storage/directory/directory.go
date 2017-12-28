package directory

import (
	"os"

	"github.com/anothermemory/lib/pkg/storage"
	"github.com/anothermemory/lib/pkg/unit"
	"github.com/pkg/errors"
)

type directoryStorage struct {
	rootDir string
}

// NewDirectoryStorage creates new storage which uses filesystem to store units
func NewDirectoryStorage(rootDir string) storage.Storage {
	return &directoryStorage{rootDir: rootDir}
}

func (s *directoryStorage) RootDir() string {
	return s.rootDir
}

func (s *directoryStorage) SaveUnit(u unit.Unit) error {
	if !s.IsCreated() {
		return errors.New("storage is not created yet and cannot be used")
	}
	return newPersistentUnitFromUnit(s.rootDir, u, s).save()
}

func (s *directoryStorage) RemoveUnit(u unit.Unit) error {
	if !s.IsCreated() {
		return errors.New("storage is not created yet and cannot be used")
	}

	return newPersistentUnitFromUnit(s.rootDir, u, s).remove()
}

func (s *directoryStorage) LoadUnit(id string) (unit.Unit, error) {
	if !s.IsCreated() {
		return nil, errors.New("storage is not created yet and cannot be used")
	}
	return newPersistentUnitFromID(s.rootDir, id, s).load()
}

func (s *directoryStorage) IsCreated() bool {
	_, err := os.Stat(s.rootDir)

	return err == nil
}

func (s *directoryStorage) Create() error {
	return errors.Wrap(os.MkdirAll(s.rootDir, os.ModePerm), "failed to create storage")
}

func (s *directoryStorage) Remove() error {
	return errors.Wrap(os.RemoveAll(s.rootDir), "failed to remove storage")
}
