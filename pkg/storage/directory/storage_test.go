package directory_test

import (
	"testing"

	"github.com/anothermemory/lib/pkg/storage"
	"github.com/anothermemory/lib/pkg/storage/directory"
	"github.com/anothermemory/lib/pkg/storage/test"
)

func TestDirectoryInMemoryStorage(t *testing.T) {
	test.RunStorageTests(t, createDirectoryInMemoryStorage)
}

func createDirectoryInMemoryStorage() storage.Storage {
	return directory.NewDirectoryInMemoryStorage()
}
