package classpath

import (
	"archive/zip"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"path/filepath"
)

type ZipEntry struct {
	absDir string
	zipRc  *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		log.Println("Error getting absolute path:", err)
	}
	return &ZipEntry{absDir, nil}
}

func (z *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	if z.zipRc == nil {
		err := z.openJar()
		if err != nil {
			return nil, z, err
		}
	}
	classFile := z.findClass(className)
	if classFile == nil {
		return nil, nil, errors.Wrap(ERR_CLASS_NOT_FOUND, fmt.Sprintf("class name:%s", className))
	}
	data, err := readClass(classFile)
	if err != nil {
		return nil, nil, err
	}
	return data, z, nil
}

func (z *ZipEntry) String() string {
	return z.absDir
}

func (z *ZipEntry) openJar() error {
	r, err := zip.OpenReader(z.absDir)
	if err == nil {
		z.zipRc = r
	}
	return err
}

func (z *ZipEntry) findClass(className string) *zip.File {
	for _, f := range z.zipRc.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(file *zip.File) ([]byte, error) {
	rc, err := file.Open()
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}
