package directory

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/anothermemory/lib/pkg/storage"
	"github.com/anothermemory/lib/pkg/unit"
	"github.com/pkg/errors"
)

type persistentUnit struct {
	unit     unit.Unit
	location location
	storage  storage.Storage
}

func newPersistentUnitFromUnit(rootDir string, u unit.Unit, s storage.Storage) *persistentUnit {
	return &persistentUnit{unit: u, location: *newLocation(rootDir, u.ID()), storage: s}
}

func newPersistentUnitFromID(rootDir string, id string, s storage.Storage) *persistentUnit {
	return &persistentUnit{unit: nil, location: *newLocation(rootDir, id), storage: s}
}

func (p *persistentUnit) marshalUnit(u unit.Unit) ([]byte, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (p *persistentUnit) marshalListUnit(u unit.List) ([]byte, error) {
	marshalItemsOriginal := u.MarshalItems()
	disableListItemsMarshal(u)

	bytes, err := p.marshalUnit(u)
	u.SetMarshalItems(marshalItemsOriginal)

	return bytes, err
}

func (p *persistentUnit) save() error {
	if nil == p.unit {
		return errors.New("cannot operate on nil unit")
	}
	err := os.MkdirAll(p.location.dirPath, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to create required directories")
	}

	var bytes []byte

	switch s := p.unit.(type) {
	case unit.List:
		bytes, err = p.marshalListUnit(s)
	default:
		bytes, err = p.marshalUnit(s)
	}

	if err != nil {
		return errors.Wrap(err, "failed to marshall unit")
	}
	return errors.Wrap(ioutil.WriteFile(p.location.fullPath, bytes, os.ModePerm), "failed to write file")
}

func (p *persistentUnit) remove() error {
	return errors.Wrap(os.Remove(p.location.fullPath), "failed to remove unit")
}

func disableListItemsMarshal(u unit.Unit) {
	if s, isList := u.(unit.List); isList {
		s.SetMarshalItems(false)
	}
}

func (p *persistentUnit) loadListItems(u unit.Unit) error {
	var items []unit.Unit
	if s, isList := u.(unit.List); isList {
		for _, ui := range s.Items() {
			i, err := p.storage.LoadUnit(ui.ID())
			if err != nil {
				return errors.Wrapf(err, "Failed to load list unit with ID: %s", ui.ID())
			}
			items = append(items, i)
		}
		s.SetItems(items)
	}
	return nil
}

type persistentUnitJSON struct {
	Type unit.Type `json:"type"`
}

func (p *persistentUnit) load() (unit.Unit, error) {
	data, err := ioutil.ReadFile(p.location.fullPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read unit")
	}

	var uj persistentUnitJSON
	err = json.Unmarshal(data, &uj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal unit type")
	}

	u := uj.Type.NewObject()
	disableListItemsMarshal(u)

	err = json.Unmarshal(data, &u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal unit")
	}

	err = p.loadListItems(u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load list unit items")
	}

	return u, nil
}
