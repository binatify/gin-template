package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type tempCfg struct {
	Name string `json:"name"`
}

func TestLoad(t *testing.T) {
	assertion := assert.New(t)

	var cfg tempCfg

	// for existed file
	{
		err := Load("test", "./", &cfg)
		assertion.Nil(err)
		assertion.Equal(cfg.Name, "config_test")
	}

	// auto try application.json
	{
		err := Load("not exist", "./", &cfg)
		assertion.Nil(err)
		assertion.Equal(cfg.Name, "config")
	}

	// try not exist files
	{
		err := Load("not exist", "./ddd/", &cfg)
		assertion.NotNil(err)
	}
}
