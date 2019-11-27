package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	rootPath, appName, moduleName string

	nameRe   = regexp.MustCompile("gin_template")
	moduleRe = regexp.MustCompile("github.com/binatify/gin-template/example")
)

func init() {
	flag.StringVar(&moduleName, "module", "", "gin-template -module=github.com/binatify/gin-template/example new example")
}

func main() {
	rootPath, _ = os.Getwd()
	parseArgs()
	copyTemplate()
	renderingTemplate()
}

func parseArgs() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 || args[0] != "new" || args[1] == "" {
		panic("invalid args, use `gin-template new hello` or `-h` to show details.")
	}

	appName = args[1]
	if moduleName == "" {
		moduleName = appName
	}
}

func copyTemplate() {
	tempFolder := "_gin-template"

	cloneCmd := exec.Command("git", "clone", "git@github.com:binatify/gin-template.git", tempFolder)
	if err := cloneCmd.Run(); err != nil {
		panic(err)
	}

	copyCmd := exec.Command("cp", "-rf", tempFolder+"/example", appName)
	if err := copyCmd.Run(); err != nil {
		panic(err)
	}

	cloneClear := exec.Command("rm", "-rf", tempFolder)
	if err := cloneClear.Run(); err != nil {
		panic(err)
	}
}

func renderingTemplate() {
	err := filepath.Walk(rootPath+"/"+appName, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return renderFile(path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func renderFile(path string) error {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var replacedContent []byte

	if isUpdateNameFile(path) {
		replacedContent = nameRe.ReplaceAll(fileContent, []byte(appName))
	} else {
		replacedContent = moduleRe.ReplaceAll(fileContent, []byte(moduleName))
	}

	return ioutil.WriteFile(path, replacedContent, 0666)
}

func isUpdateNameFile(path string) bool {
	return strings.Contains(path, "config/application") || strings.Contains(path, "Makefile")
}
