package directory

import (
	"encoding/json"
	"os"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/pkg/errors"
)

type persistentUnit struct {
	unit     unit.Unit
	location location
	storage  persistentUnitStorage
}

func newPersistentUnit(u unit.Unit, l location, s persistentUnitStorage) *persistentUnit {
	return &persistentUnit{unit: u, location: l, storage: s}
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
	err := p.storage.mkdirAll(p.location.dirPath, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to create required directories")
	}

	var bytes []byte

	switch p.unit.Type() {
	case unit.TypeList:
		bytes, err = p.marshalListUnit(p.unit.(unit.List))
	default:
		bytes, err = p.marshalUnit(p.unit)
	}

	if err != nil {
		return errors.Wrap(err, "failed to marshall unit")
	}

	return errors.Wrap(p.storage.writeFile(p.location.fullPath, bytes, os.ModePerm), "failed to write file")
}

func (p *persistentUnit) remove() error {
	return errors.Wrap(p.storage.removeDir(p.location.fullPath), "failed to remove unit")
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
	data, err := p.storage.readFile(p.location.fullPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read unit")
	}

	var uj persistentUnitJSON
	err = json.Unmarshal(data, &uj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal unit type")
	}

	if uj.Type == unit.TypeList {
		return p.loadList(data)
	}
	return p.loadNotList(uj.Type, data)
}

func (p *persistentUnit) loadList(data []byte) (unit.Unit, error) {
	baseList := unit.NewList()
	err := json.Unmarshal(data, &baseList)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal list unit base data to get updated value")
	}

	u := unit.NewList(unit.OptionListMarshalItems(false), unit.OptionClockMockPartial(baseList.Updated(), baseList.Updated()))
	err = json.Unmarshal(data, &u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal list unit")
	}

	err = p.loadListItems(u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load list unit items")
	}

	return u, nil
}

func (p *persistentUnit) loadNotList(t unit.Type, data []byte) (unit.Unit, error) {
	u := t.NewObject()
	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal unit")
	}
	return u, nil
}
