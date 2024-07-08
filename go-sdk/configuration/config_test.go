package configuration

import (
	"testing"

	"github.com/Ibrahimsyah/smolf/go-sdk/env"
	"github.com/stretchr/testify/assert"
)

type user struct {
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
}

type config struct {
	Hello string `yaml:"hello"`
	User  user   `yaml:"user"`
}

func TestReadConfig_SuccessDevelopment(t *testing.T) {
	var result config
	err := ReadConfig(ReadConfigParam{
		Dest:     &result,
		Folder:   "config",
		FileName: "configuration",
		Env:      "development",
	})
	assert.NoError(t, err)
	assert.Equal(t, config{
		Hello: "world",
		User: user{
			FirstName: "john",
			LastName:  "doe",
		},
	}, result)
}

func TestReadConfig_SuccessProduction(t *testing.T) {
	var result config
	err := ReadConfig(ReadConfigParam{
		Dest:     &result,
		Folder:   "config",
		FileName: "configuration",
		Env:      env.PRODUCTION,
	})
	assert.NoError(t, err)
	assert.Equal(t, config{
		Hello: "world-2",
		User: user{
			FirstName: "john",
			LastName:  "wick",
		},
	}, result)
}
