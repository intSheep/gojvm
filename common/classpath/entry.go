package classpath

import (
	"os"
	"strings"
)

var pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	ReadClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	switch {
	case strings.Contains(path, pathListSeparator):
		return newCompositeEntry(path)
	case strings.HasPrefix(path, "*"):
		return newWildcardEntry(path)
	case strings.HasPrefix(path, ".jar"), strings.HasPrefix(path, ".JAR"), strings.HasPrefix(path, ".zip"), strings.HasPrefix(path, ".ZIP"):
		return newZipEntry(path)
	default:
		return newDirEntry(path)
	}
}
