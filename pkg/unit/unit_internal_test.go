package unit

import (
	"testing"

	"encoding/json"

	"fmt"
	"strings"
	"time"

	"github.com/anothermemory/lib/pkg/clock"
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

func jsonUnit(u Unit) string {
	return jsonUnitRaw(u.ID(), u.Title(), u.Type(), u.Created(), u.Updated())
}

func jsonUnitRaw(id, title string, unitType Type, created, updated time.Time) string {
	return fmt.Sprintf(formatJSONUnit, id, title, unitType, jsonTime(created), jsonTime(updated))
}

func jsonUnitDummy(unitType Type) string {
	return jsonUnitRaw("123", "MyUnit", unitType, createdTime, updatedTime)
}

func freezeUnitTime(u *baseUnit, created, updated time.Time) {
	u.clock = clock.Freeze(created)
	u.refreshCreated()

	u.clock = clock.Freeze(updated)
	u.refreshUpdated()
}

func TestNewUnit_Clock(t *testing.T) {
	u := newBaseUnit(TypeUnit)
	freezeUnitTime(u, createdTime, createdTime)

	assert.Equal(t, createdTime, u.Created())
	assert.Equal(t, createdTime, u.Updated())

	u.clock = clock.Freeze(updatedTime)
	u.refreshUpdated()

	assert.Equal(t, createdTime, u.Created())
	assert.Equal(t, updatedTime, u.Updated())
}

func TestBaseUnit_MarshalJSON(t *testing.T) {
	u := newBaseUnit(TypeUnit)
	u.SetTitle("MyUnit")
	freezeUnitTime(u, createdTime, updatedTime)

	bytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.JSONEq(t, toJSON(jsonUnit(u)), string(bytes))
}

func TestBaseUnit_UnmarshalJSON(t *testing.T) {
	u := NewUnit()

	err := json.Unmarshal([]byte(toJSON(jsonUnitDummy(TypeUnit))), &u)
	assert.NoError(t, err)
	assert.Equal(t, "123", u.ID())
	assert.Equal(t, "MyUnit", u.Title())
	assert.Equal(t, TypeUnit, u.Type())
	assert.Equal(t, createdTime, u.Created())
	assert.Equal(t, updatedTime, u.Updated())
}
