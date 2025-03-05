package viper

import (
	"github.com/spf13/viper"
	"gogo/backend/configs"
	"gogo/backend/global"
	"gogo/cmd/server/conf"
	"gopkg.in/yaml.v3"
)

func Init() {

	v := viper.NewWithOptions()
	v.SetConfigType("yaml")
	config := configs.ServerConfig{}

	if err := yaml.Unmarshal(conf.AppYaml, &config); err != nil {
		panic(err)
	}

	global.Viper = v
}
