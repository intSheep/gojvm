package classpath

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var entries []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entries = append(entries, newEntry(path))
	}
	return entries
}

func (cp CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range cp {
		data, from, err := entry.ReadClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.Wrap(ERR_CLASS_NOT_FOUND, fmt.Sprintf("class name:%s", className))
}

func (cp CompositeEntry) String() string {
	var sb strings.Builder
	for _, entry := range cp {
		sb.WriteString(entry.String())
		sb.WriteString(pathListSeparator)
	}
	return sb.String()
}
