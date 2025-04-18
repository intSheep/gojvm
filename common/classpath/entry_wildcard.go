package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	entries := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		ext := filepath.Ext(path)
		if strings.HasSuffix(strings.ToLower(ext), ".jar") {
			jarEntry := newZipEntry(path)
			entries = append(entries, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return entries
}
