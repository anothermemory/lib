package unit

import (
	"encoding/json"
	"time"

	"github.com/anothermemory/lib/pkg/clock"
	"github.com/anothermemory/lib/pkg/idgen"
	"github.com/pkg/errors"
)

// Unit represents simplest unit which actually does nothing but used as a base for all other units
type Unit interface {
	json.Marshaler
	json.Unmarshaler
	ID() string
	Title() string
	SetTitle(title string)
	Type() Type
	Created() time.Time
	Updated() time.Time
}

// baseUnit represents default implementation of Unit interface
type baseUnit struct {
	Unit
	id          string
	title       string
	unitType    Type
	created     time.Time
	updated     time.Time
	clock       clock.Clock
	idGenerator idgen.Generator
}

// NewUnit creates new Unit with given title. Unit id is generated automatically
func NewUnit(options ...func(u interface{})) Unit {
	u := newBaseUnit(TypeUnit)
	initUnitOptions(u, options...)
	initUnit(u)

	return u
}

func newBaseUnit(unitType Type) *baseUnit {
	return &baseUnit{idGenerator: idgen.NewUUID(), title: "", unitType: unitType, clock: clock.NewReal()}
}

func initUnit(u *baseUnit) {
	u.id = u.idGenerator.Generate()
	t := u.clock.Now()
	u.created = t
	u.updated = t
}

func initUnitOptions(u interface{}, options ...func(u interface{})) {
	for _, option := range options {
		option(u)
	}
}

// ID returns unit id
func (u *baseUnit) ID() string {
	return u.id
}

// Title returns unit title
func (u *baseUnit) Title() string {
	return u.title
}

// SetTitle sets new unit title
func (u *baseUnit) SetTitle(title string) {
	u.title = title
	u.refreshUpdated()
}

func (u *baseUnit) Type() Type {
	return u.unitType
}

func (u *baseUnit) Created() time.Time {
	return u.created
}

func (u *baseUnit) Updated() time.Time {
	return u.updated
}

func (u *baseUnit) refreshUpdated() {
	u.updated = u.clock.Now()
}

type baseUnitJSON struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Type    Type      `json:"type"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func (u *baseUnit) fromJSONStruct(j baseUnitJSON) error {
	if j.Type != u.Type() {
		return errors.Errorf("Unexpected unit type in json data received. Expected: %s, Actual: %s", u.Type(), j.Type)
	}
	u.id = j.ID
	u.created = j.Created
	u.updated = j.Updated
	u.title = j.Title

	return nil
}

func (u *baseUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(baseUnitJSON{ID: u.id, Title: u.title, Type: u.unitType, Created: u.created, Updated: u.updated})
}

func (u *baseUnit) UnmarshalJSON(b []byte) error {
	var jsonData baseUnitJSON
	err := json.Unmarshal(b, &jsonData)

	if err != nil {
		return err
	}

	return u.fromJSONStruct(jsonData)
}

// OptionIDGeneratorUUID is an option that sets internal UUID generator for a unit to UUID implementation
func OptionIDGeneratorUUID() func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseUnit); converted {
			o.idGenerator = idgen.NewUUID()
		}
	}
}

// OptionIDGeneratorMock is an option that sets internal UUID generator for a unit to return same value each time
func OptionIDGeneratorMock(id string) func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseUnit); converted {
			o.idGenerator = idgen.NewMock(id)
		}
	}
}

// OptionClockReal is an option that sets internal clock for a unit to return real time
func OptionClockReal() func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseUnit); converted {
			o.clock = clock.NewReal()
		}
	}
}

// OptionClockMock is an option that sets internal clock for a unit to return mocked values
func OptionClockMock(t ...time.Time) func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseUnit); converted {
			o.clock = clock.NewMock(t...)
		}
	}
}

// OptionClockMockPartial is an option that sets internal clock for a unit to return mocked value once
func OptionClockMockPartial(t ...time.Time) func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseUnit); converted {
			o.clock = clock.NewMockPartial(t...)
		}
	}
}

// OptionTitle is an option that sets title for a unit to the provided value
func OptionTitle(t string) func(u interface{}) {
	return func(u interface{}) {
		if o, converted := u.(*baseUnit); converted {
			o.title = t
		}
	}
}
