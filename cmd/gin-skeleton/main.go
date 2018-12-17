package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	copySkeleton()
	renderingSkeleton()
}

func copySkeleton() {
	rmCmd := exec.Command("rm", "-rf", "../../skeleton")
	if err := rmCmd.Run(); err != nil {
		panic(err)
	}

	cloneCmd := exec.Command("cp", "-rf", "../../example", "../../skeleton")
	if err := cloneCmd.Run(); err != nil {
		panic(err)
	}
}

func renderingSkeleton() {
	err := filepath.Walk("../../skeleton", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return renderFile(path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

var (
	moduleRe = regexp.MustCompile("github.com/binatify/gin-template/example")
	nameRe   = regexp.MustCompile("gin_template")
)

func renderFile(path string) error {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var replacedConent []byte

	if isUpdateNameFile(path) {
		replacedConent = nameRe.ReplaceAll(fileContent, []byte("{{.Name}}"))
	} else {
		replacedConent = moduleRe.ReplaceAll(fileContent, []byte("{{.Module}}"))
	}

	return ioutil.WriteFile(path, replacedConent, 0666)
}

func isUpdateNameFile(path string) bool {
	return strings.Contains(path, "config/application") || strings.Contains(path, "Makefile")
}
