package classpath

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

var DefaultJre = "./jre"

var ERR_JRE_NOT_FOUND = errors.New("jre not found")

type Classpath struct {
	bootClasspath Entry //启动类路径，包括JVM标准库
	extClasspath  Entry //拓展类路径，通常位于jre/lib/ext目录或者由java.ext.dirs系统属性指定的目录
	userClasspath Entry //最常用的类路径，使用-cp或-classpath命令行选项来指定
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	if !strings.HasSuffix(className, ".class") {
		className = className + ".class"
	}
	if data, entry, err := c.bootClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := c.extClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}
	return c.userClasspath.ReadClass(className)
}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}

func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir, err := getJreDir(jreOption)
	if err != nil {
		logrus.Infof("parseBootAndExtClasspath failed: %v", err)
	}
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}
func getJreDir(jreOption string) (string, error) {
	if jreOption != "" && isPathExit(jreOption) {
		return jreOption, nil
	}
	if isPathExit(DefaultJre) {
		return DefaultJre, nil
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre"), nil
	}
	return "", ERR_JRE_NOT_FOUND
}

func isPathExit(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
