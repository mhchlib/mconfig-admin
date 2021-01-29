package config

import (
	log "github.com/mhchlib/logger"
	"github.com/spf13/viper"
)

func Init() {
	viper.AddConfigPath("conf/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal("Fatal error config file: %s \n", err)
	}
}
