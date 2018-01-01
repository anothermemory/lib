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

func (s *directoryStorage) mkdirAll(path string, perm os.FileMode) error {
	return s.fs.MkdirAll(path, perm)
}

func (s *directoryStorage) removeDir(name string) error {
	return s.fs.RemoveAll(name)
}

func (s *directoryStorage) writeFile(filename string, data []byte, perm os.FileMode) error {
	return s.fsUtil.WriteFile(filename, data, perm)
}

func (s *directoryStorage) readFile(filename string) ([]byte, error) {
	return s.fsUtil.ReadFile(filename)
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
	return newPersistentUnit(u, *newLocation(s.RootDir(), u.ID()), s).save()
}

func (s *directoryStorage) RemoveUnit(u unit.Unit) error {
	if !s.IsCreated() {
		return errors.New("storage is not created yet and cannot be used")
	}
	if nil == u {
		return errors.New("cannot operate on nil unit")
	}
	return newPersistentUnit(u, *newLocation(s.RootDir(), u.ID()), s).remove()
}

func (s *directoryStorage) LoadUnit(id string) (unit.Unit, error) {
	if !s.IsCreated() {
		return nil, errors.New("storage is not created yet and cannot be used")
	}
	if len(id) == 0 {
		return nil, errors.New("cannot operate on nil unit")
	}
	return newPersistentUnit(nil, *newLocation(s.RootDir(), id), s).load()
}

func (s *directoryStorage) IsCreated() bool {
	_, err := s.fs.Stat(s.rootDir)

	return err == nil
}

func (s *directoryStorage) Create() error {
	return errors.Wrap(s.mkdirAll(s.rootDir, os.ModePerm), "failed to create storage")
}

func (s *directoryStorage) Remove() error {
	return errors.Wrap(s.removeDir(s.rootDir), "failed to remove storage")
}
