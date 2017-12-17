package storage

import (
	"testing"

	"github.com/anothermemory/lib/pkg/unit"
	"github.com/stretchr/testify/assert"
)

func TestPersistentUnit(t *testing.T) {
	const rootPath = "/app/data"
	cases := []struct{ id, expectedDir, expectedName, expectedPath string }{
		{"123", "/app/data/40", "bd001563085fc35165329ea1ff5c5ecbdbbeef.json", "/app/data/40/bd001563085fc35165329ea1ff5c5ecbdbbeef.json"},
		{"5a044285-0d10-44db-b233-d5c6898488da", "/app/data/d2", "68b2cbf0d664f452b8ef139a5581c1a087f487.json", "/app/data/d2/68b2cbf0d664f452b8ef139a5581c1a087f487.json"},
	}

	for _, tc := range cases {
		u := unit.NewUnit("A")
		u.SetID(tc.id)

		p := newPersistentUnit(rootPath, u)
		assert.Equal(t, tc.expectedDir, p.Directory())
		assert.Equal(t, tc.expectedName, p.Filename())
		assert.Equal(t, tc.expectedPath, p.Path())
	}
}

func TestPersistentUnit_SameId(t *testing.T) {
	const id = "123"
	const rootPath = "/app/data"
	u1 := unit.NewUnit("A")
	u1.SetID(id)
	p1 := newPersistentUnit(rootPath, u1)

	u2 := unit.NewUnit("B")
	u2.SetID(id)
	p2 := newPersistentUnit(rootPath, u2)

	assert.Equal(t, p1.Directory(), p2.Directory())
	assert.Equal(t, p1.Filename(), p2.Filename())
	assert.Equal(t, p1.Path(), p2.Path())
}
