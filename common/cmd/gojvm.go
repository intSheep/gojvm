package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gojvm/common/classfile"
	"gojvm/common/classpath"
	"log"
	"os"
	"strings"
)

var Version = "Version 1.0.0"

func startJVM() *cli.App {
	return &cli.App{
		Name:  "gojvm",
		Usage: "A Jvm implement by Go",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "classpath",
				Aliases: []string{"cp"},
				Usage:   "Set the classpath",
			},
			&cli.BoolFlag{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Print the version",
			},
			&cli.StringFlag{
				Name:    "Xjre",
				Aliases: []string{"xjre"},
				Usage:   "Set the path to jre",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("version") {
				fmt.Println(Version)
				return nil
			}

			if c.NArg() < 1 {
				return cli.Exit("Error: No class specified", 1)
			}

			className := c.Args().First()
			classPath := c.String("classpath")
			xjrePath := c.String("Xjre")
			args := c.Args().Tail()
			cp := classpath.Parse(xjrePath, classPath)
			fmt.Printf("ClassPath:%s class:%s args:%s\n", cp, className, args)
			newClassName := strings.Replace(className, ".", "/", -1)
			cf, err := loadClass(newClassName, cp)
			if err != nil {
				return err
			}
			fmt.Println("===============================================================================")
			cf.Show()
			return nil
		},
	}
}

func loadClass(className string, cp *classpath.Classpath) (*classfile.ClassFile, error) {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		return nil, errors.Errorf("Counld not find or load main class %s\n", className)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		return nil, err
	}
	return cf, nil
}

func main() {
	app := startJVM()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
