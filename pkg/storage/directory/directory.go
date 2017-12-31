package directory

import (
	"os"

	"github.com/anothermemory/lib/pkg/storage"
	"github.com/anothermemory/lib/pkg/unit"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type directoryStorage struct {
	rootDir string
	fs      afero.Fs
	fsUtil  *afero.Afero
}

// NewDirectoryStorage creates new storage which uses filesystem to store units
func NewDirectoryStorage(rootDir string) storage.Storage {
	fs := afero.NewOsFs()
	return &directoryStorage{rootDir: rootDir, fs: fs, fsUtil: &afero.Afero{Fs: fs}}
}

// NewDirectoryInMemoryStorage creates new storage which uses memory to store units
func NewDirectoryInMemoryStorage() storage.Storage {
	fs := afero.NewMemMapFs()
	return &directoryStorage{rootDir: "/anothermemory", fs: fs, fsUtil: &afero.Afero{Fs: fs}}
}

func (s *directoryStorage) RootDir() string {
	return s.rootDir
}

func (s *directoryStorage) SaveUnit(u unit.Unit) error {
	if !s.IsCreated() {
		return errors.New("storage is not created yet and cannot be used")
	}
	if nil == u {
		return errors.New("cannot operate on nil unit")
	}
	return newPersistentUnitFromUnit(s.rootDir, u, s).save()
}

func (s *directoryStorage) RemoveUnit(u unit.Unit) error {
	if !s.IsCreated() {
		return errors.New("storage is not created yet and cannot be used")
	}
	if nil == u {
		return errors.New("cannot operate on nil unit")
	}
	return newPersistentUnitFromUnit(s.rootDir, u, s).remove()
}

func (s *directoryStorage) LoadUnit(id string) (unit.Unit, error) {
	if !s.IsCreated() {
		return nil, errors.New("storage is not created yet and cannot be used")
	}
	if len(id) == 0 {
		return nil, errors.New("cannot operate on nil unit")
	}
	return newPersistentUnitFromID(s.rootDir, id, s).load()
}

func (s *directoryStorage) IsCreated() bool {
	_, err := s.fs.Stat(s.rootDir)

	return err == nil
}

func (s *directoryStorage) Create() error {
	return errors.Wrap(s.fs.MkdirAll(s.rootDir, os.ModePerm), "failed to create storage")
}

func (s *directoryStorage) Remove() error {
	return errors.Wrap(s.fs.RemoveAll(s.rootDir), "failed to remove storage")
}
