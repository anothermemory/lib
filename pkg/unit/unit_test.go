package unit_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

var createdTime = time.Date(2017, 11, 24, 17, 0, 0, 0, time.Local)
var updatedTime = time.Date(2017, 12, 25, 18, 0, 0, 0, time.Local)

const formatJSONString = `{%s}`
const formatJSONUnit = `"id": "%s", "title": "%s", "type": "%s", "created":"%s", "updated":"%s"`

func toJSON(items ...string) string {
	return fmt.Sprintf(formatJSONString, strings.Join(items, ","))
}
func jsonTime(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}

func jsonUnit(u unit.Unit) string {
	return jsonUnitRaw(u.ID(), u.Title(), u.Type(), u.Created(), u.Updated())
}

func jsonUnitRaw(id, title string, unitType unit.Type, created, updated time.Time) string {
	return fmt.Sprintf(formatJSONUnit, id, title, unitType, jsonTime(created), jsonTime(updated))
}

func jsonUnitDummy(unitType unit.Type) string {
	return jsonUnitRaw("123", "MyUnit", unitType, createdTime, updatedTime)
}

func TestNewUnit(t *testing.T) {
	u := unit.NewUnit()
	assert.Equal(t, unit.TypeUnit, u.Type())
	assert.NotNil(t, u.ID())
	assert.NotNil(t, u.Created())
	assert.NotNil(t, u.Updated())
}

func TestNewUnit_MockID(t *testing.T) {
	const id = "123"
	u := unit.NewUnit(unit.OptionIDGeneratorMock(id))
	assert.Equal(t, id, u.ID())
}
func TestNewUnit_Title(t *testing.T) {
	const title = "title"
	u := unit.NewUnit(unit.OptionTitle(title))
	assert.Equal(t, title, u.Title())
}

func TestBaseUnit_ID(t *testing.T) {
	assert.NotEqual(t, unit.NewUnit().ID(), unit.NewUnit().ID())
}

func TestBaseUnit_Title(t *testing.T) {
	u := unit.NewUnit()
	u.SetTitle("MyUnit")
	assert.Equal(t, "MyUnit", u.Title())
}

func TestBaseUnit_Created(t *testing.T) {
	u := unit.NewUnit(unit.OptionClockMock(createdTime, updatedTime))

	assert.Equal(t, createdTime, u.Created())

	u.SetTitle("New title will not change created time")

	assert.Equal(t, createdTime, u.Created())
}

func TestBaseUnit_Updated(t *testing.T) {
	u := unit.NewUnit(unit.OptionClockMock(createdTime, updatedTime))

	assert.Equal(t, createdTime, u.Updated())

	u.SetTitle("New title will change updated time")

	assert.Equal(t, updatedTime, u.Updated())
}

func TestBaseUnit_MarshalJSON(t *testing.T) {
	u := unit.NewUnit(unit.OptionClockMock(createdTime, updatedTime))
	u.SetTitle("MyUnit")

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, toJSON(jsonUnit(u)), string(bytes))
}

func TestBaseUnit_UnmarshalJSON(t *testing.T) {
	u := unit.NewUnit()

	err := json.Unmarshal([]byte(toJSON(jsonUnitDummy(unit.TypeUnit))), &u)
	assert.NoError(t, err)
	assert.Equal(t, "123", u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, unit.TypeUnit, u.Type())
	assert.Equal(t, createdTime, u.Created())
	assert.Equal(t, updatedTime, u.Updated())
}
