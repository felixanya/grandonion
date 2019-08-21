package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"reflect"
)

func main() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./examples/loadConf/toml-config.toml")

	viper.SetConfigType("toml")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config.toml file: %s \n", err.Error()))
	} else {
		fmt.Println("Using config.toml:", viper.ConfigFileUsed())
	}

	configMap := viper.Get(".")

	writeDbConf := fmt.Sprintf("%s.database.write_db", "prod")
	serviceMap := viper.GetStringMap(writeDbConf)

	fmt.Println(">>>service map>>>:", serviceMap)
	fmt.Println(">>>config.toml map>>>:", configMap)

	fmt.Println(">>>db port>>>:", serviceMap["port"])
	fmt.Println(">>>if port equals 3306>>>:", serviceMap["port"] == int64(3306))
	fmt.Println(">>>port data type>>>:", reflect.TypeOf(serviceMap["port"]).String())

	//select {}
}
