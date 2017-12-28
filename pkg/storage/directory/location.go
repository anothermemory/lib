package directory

import (
	"crypto/sha1"
	"fmt"
	"path/filepath"
)

type location struct {
	dirPath  string
	filename string
	fullPath string
}

func newLocation(rootDir string, id string) *location {
	data := sha1.Sum([]byte(id))
	hash := fmt.Sprintf("%x", data)
	d := filepath.Join(rootDir, hash[0:2])
	fn := fmt.Sprintf("%s.json", hash[2:40])
	path := filepath.Join(d, fn)

	return &location{dirPath: d, filename: fn, fullPath: path}
}
