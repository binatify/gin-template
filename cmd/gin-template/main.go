package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type AppInfo struct {
	Name     string
	Module   string
	Template string
}

var appInfo AppInfo

func init() {
	flag.StringVar(&appInfo.Name, "name", "gin-template-example", "gin-template -name=gin-template-example")
	flag.StringVar(&appInfo.Module, "module", "github.com/binatify/gin-template/example", "gin-template -module=github.com/binatify/gin-template/example")
}

func main() {
	flag.Parse()
	copyTemplate()
	renderingTemplate()
}

func copyTemplate() {
	tempFolder := "_gin-template"

	cloneCmd := exec.Command("git", "clone", "git@github.com:binatify/gin-template.git", tempFolder)
	if err := cloneCmd.Run(); err != nil {
		panic(err)
	}

	copyCmd := exec.Command("cp", "-rf", tempFolder+"/skeleton", appInfo.Name)
	if err := copyCmd.Run(); err != nil {
		panic(err)
	}

	cloneClear := exec.Command("rm", "-rf", tempFolder)
	if err := cloneClear.Run(); err != nil {
		panic(err)
	}
}

func renderingTemplate() {
	root, _ := os.Getwd()
	err := filepath.Walk(root+"/"+appInfo.Name, func(path string, info os.FileInfo, err error) error {
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

	t := template.Must(template.New(path).Parse(string(fileContent)))

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	return t.Execute(file, &appInfo)
}
