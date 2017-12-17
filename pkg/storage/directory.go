package storage

import (
	"crypto/sha1"
	"path/filepath"

	"fmt"

	"os"

	"encoding/json"
	"io/ioutil"

	"github.com/anothermemory/lib/pkg/unit"
)

type directoryStorage struct {
	rootDir string
}

// NewDirectoryStorage creates new storage which uses filesystem to store units
func NewDirectoryStorage(rootDir string) Storage {
	return &directoryStorage{rootDir: rootDir}
}

//
//func NewDirectoryStorageFromJsonConfig(config json.RawMessage) (Storage, error) {
//	var result directoryStorage
//	err := json.Unmarshal(config, &result)
//	if err != nil {
//		return nil, &JsonConfigError{details: err.Error()}
//	}
//	return &result, nil
//}
//
func (s *directoryStorage) SaveUnit(u unit.Unit) error {
	if !s.IsCreated() {
		return ErrNotCreated
	}
	return newPersistentUnit(s.rootDir, u).save()
}

func (s *directoryStorage) RemoveUnit(u unit.Unit) error {
	if !s.IsCreated() {
		return ErrNotCreated
	}

	return newPersistentUnit(s.rootDir, u).remove()
}

func (s *directoryStorage) LoadUnit(id string) (unit.Unit, error) {
	panic("implement me")
}

func (s *directoryStorage) LoadUnits(ids []string) (map[string]unit.Unit, error) {
	panic("implement me")
}

func (s *directoryStorage) Create() error {
	return os.MkdirAll(s.rootDir, os.ModePerm)
}

func (s *directoryStorage) IsCreated() bool {
	_, err := os.Stat(s.rootDir)

	return err == nil
}

func (s *directoryStorage) Remove() error {
	return os.RemoveAll(s.rootDir)
}

type persistentUnit struct {
	unit      unit.Unit
	directory string
	filename  string
	path      string
}

func newPersistentUnit(rootDir string, u unit.Unit) *persistentUnit {
	data := sha1.Sum([]byte(u.ID()))
	hash := fmt.Sprintf("%x", data)
	d := filepath.Join(rootDir, hash[0:2])
	fn := fmt.Sprintf("%s.json", hash[2:40])
	path := filepath.Join(d, fn)

	return &persistentUnit{unit: u, directory: d, filename: fn, path: path}
}

type persistentListJSON struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Type  string   `json:"type"`
	Items []string `json:"items"`
}

func (p *persistentUnit) marshalUnit(u unit.Unit) ([]byte, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (p *persistentUnit) marshalListUnit(u unit.List) ([]byte, error) {
	var items []string
	for _, i := range u.Items() {
		items = append(items, i.ID())
	}
	return json.Marshal(persistentListJSON{ID: u.ID(), Title: u.Title(), Type: u.Type(), Items: items})
}

func (p *persistentUnit) save() error {
	err := os.MkdirAll(p.directory, os.ModePerm)
	if err != nil {
		return err
	}

	var bytes []byte

	switch s := p.unit.(type) {
	case unit.List:
		bytes, err = p.marshalListUnit(s)
	default:
		bytes, err = p.marshalUnit(s)
	}

	if err != nil {
		return err
	}
	return ioutil.WriteFile(p.path, bytes, os.ModePerm)
}

func (p *persistentUnit) remove() error {
	return os.Remove(p.path)
}
