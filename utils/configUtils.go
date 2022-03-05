package utils

import (
	"epos_v3/configuration"
	"github.com/tkanos/gonfig"
)

func InitConfiguration(configPath string) *configuration.BaseConfig {
	conf := configuration.BaseConfig{}
	err := gonfig.GetConf(configPath, &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}

func GetSourceConfig(configPath string) *configuration.SourceConfig {
	conf := configuration.SourceConfig{}
	err := gonfig.GetConf(configPath, &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
