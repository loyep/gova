package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config type is the global config
type Config struct {
}

func init() {
	runtimeViper := viper.New()
	runtimeViper.AddConfigPath(".")
	runtimeViper.SetConfigName("config")
	runtimeViper.SetConfigType("json")
	err := runtimeViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	runtimeViper.WatchConfig()
	runtimeViper.OnConfigChange(func(e fsnotify.Event) {
	})
}
