package config

import (
	"github.com/spf13/viper"
	"log"
	"reneat-microservice-user/helpers/utils"
	"strings"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(env)
	v.AddConfigPath("../config/")
	v.AddConfigPath("./config/")
	v.AddConfigPath("config/")
	v.AddConfigPath("/app/config/")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := v.ReadInConfig()
	if err != nil {
		println("env:", env)
		utils.DebugJson(err)
		log.Fatal("Error on parsing configuration file")
	}
	config = v
}

func GetConfig() *viper.Viper {
	return config
}
