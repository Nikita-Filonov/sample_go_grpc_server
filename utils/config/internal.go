package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadConfig(configPath string, config interface{}) error {
	if configPath == `` {
		return fmt.Errorf(`no config path`)
	}

	configBytes, err := ioutil.ReadFile(configPath)

	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(configBytes, config); err != nil {
		return err
	}

	return nil
}

func NewConfig() Config {
	var config Config

	err := ReadConfig("./infrastructure/config.yml", &config)

	if err != nil {
		panic(err)
	}

	return config
}
