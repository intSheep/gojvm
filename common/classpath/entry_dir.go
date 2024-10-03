package classpath

import (
	"log"
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		log.Println("Error getting absolute path:", err)
	}
	return &DirEntry{absDir}
}

func (d *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(d.absDir, className)
	data, err := os.ReadFile(fileName)
	return data, d, err
}

func (d *DirEntry) String() string {
	return d.absDir
}
