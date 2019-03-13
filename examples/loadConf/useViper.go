package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")

	viper.SetConfigType("toml")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err.Error()))
	} else {
		fmt.Println("Using config:", viper.ConfigFileUsed())
	}

	serviceMap := viper.GetStringMap("service")

	fmt.Println(">>>service map>>>:", serviceMap)

	select {}
}
