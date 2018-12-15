package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	TimestampFormat = "2006-01-02 15:04:05.999999"
)

type Config struct {
	Output        string `json:"output"`
	Level         string `json:"level"`
	Format        string `json:"format"`
	DisableCaller bool   `json:"disable_caller"`
}

func (c *Config) IsStdout() bool {
	return c.Output == "stdout"
}

func (c *Config) IsJsonFormat() bool {
	return c.Format == "json"
}

func (c *Config) GetOutputWriter() (io.Writer, error) {
	output := c.Output

	switch output {
	case "stdout":
		return os.Stdout, nil

	case "stderr":
		return os.Stderr, nil

	default:
		if output == "null" || output == "nil" {
			output = os.DevNull
		}

		file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("Failed to open log file %s: %v", output, err)
		}

		return file, nil
	}
}

func (c *Config) GetFormater() logrus.Formatter {
	if c.IsJsonFormat() {
		return &logrus.JSONFormatter{}
	}

	return &logrus.TextFormatter{
		DisableColors:   !c.IsStdout(),
		FullTimestamp:   true,
		TimestampFormat: TimestampFormat,
	}
}
