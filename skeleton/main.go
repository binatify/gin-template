package main

import (
	"flag"
	"os"
	"path"

	"{{.Module}}/app"
	"github.com/binatify/gin-template/base/runmode"
)

var (
	runMode string
	srcPath string
)

func init() {
	flag.StringVar(&runMode, "runMode", "development", "app run -runMode=[development|test|production]")

	flag.StringVar(&srcPath, "srcPath", "", "gin-demo -srcPath=/path/to/source")
}

func main() {
	flag.Parse()

	mode := runmode.RunMode(runMode)
	if !mode.IsValid() {
		flag.PrintDefaults()
		return
	}

	if srcPath == "" {
		var err error

		srcPath, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	} else {
		srcPath = path.Clean(srcPath)
	}

	app.New(mode, srcPath).Run()
}
