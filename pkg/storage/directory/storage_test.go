package directory_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/anothermemory/lib/pkg/storage"
	"github.com/anothermemory/lib/pkg/storage/directory"
	"github.com/anothermemory/lib/pkg/storage/test"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestDirectoryInMemoryStorage(t *testing.T) {
	test.RunStorageTests(t, createDirectoryInMemoryStorage)
}

func createDirectoryInMemoryStorage() storage.Storage {
	return directory.NewDirectoryInMemoryStorage()
}

func TestDirectoryStorage(t *testing.T) {
	dir, err := ioutil.TempDir("", "storage_directory_root")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	test.RunStorageTests(t, func() storage.Storage {
		return directory.NewDirectoryStorage(path.Join(dir, uuid.NewV4().String()))
	})
}
