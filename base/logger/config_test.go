package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig_IsFileOutput(t *testing.T) {
	assertion := assert.New(t)
	cfg := &Config{
		Output: "stdout",
	}

	assertion.True(cfg.IsStdout())
	cfg.Output = "./a.log"
	assertion.False(cfg.IsStdout())
}

func TestConfig_IsJsonFormat(t *testing.T) {
	assertion := assert.New(t)
	cfg := &Config{
		Format: "text",
	}

	assertion.False(cfg.IsJsonFormat())
	cfg.Format = "json"
	assertion.True(cfg.IsJsonFormat())
}

func TestConfig_GetFormater(t *testing.T) {
	assertion := assert.New(t)

	cfg := &Config{
		Format: "text",
	}
	formatter := cfg.GetFormater()
	assertion.NotNil(formatter)
	_, ok := formatter.(*logrus.TextFormatter)
	assertion.True(ok)

	cfg.Format = "json"
	formatter = cfg.GetFormater()
	assertion.NotNil(formatter)
	_, ok = formatter.(*logrus.JSONFormatter)
	assertion.True(ok)
}

func TestConfig_GetOutputWriter(t *testing.T) {
	assertion := assert.New(t)
	cfg := &Config{
		Output: "stdout",
	}

	out, err := cfg.GetOutputWriter()
	assertion.Nil(err)
	assertion.NotNil(out)

	// can not create folder
	cfg.Output = "./ddd/a.log"
	out, err = cfg.GetOutputWriter()
	assertion.NotNil(err)
	assertion.Nil(out)

	// can not create folder
	cfg.Output = "./a.log"
	defer os.Remove(cfg.Output)
	out, err = cfg.GetOutputWriter()
	assertion.Nil(err)
	assertion.NotNil(out)
}
