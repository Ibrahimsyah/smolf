package configuration

import (
	"fmt"
	"os"

	"github.com/Ibrahimsyah/smolf/go-sdk/env"
	"gopkg.in/yaml.v3"
)

type ReadConfigParam struct {
	Dest     interface{}
	Folder   string
	FileName string
	Env      env.ENV
}

func ReadConfig(param ReadConfigParam) error {
	if param.Env == "" {
		param.Env = env.DEVELOPMENT
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s/%s.%s.yaml", workingDir, param.Folder, param.FileName, param.Env)
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamlFile, param.Dest); err != nil {
		return err
	}

	return nil
}
