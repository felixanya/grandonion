package main

import (
    "fmt"
    "github.com/spf13/viper"
    "log"
)

type ServiceConfig struct {
    Bind        string
    Port        int

    Loggers     Loggers
}

type Loggers struct {
    Path        string
    Level       string
}

func main() {
    var config ServiceConfig

    viper.SetConfigName("viper_conf") // 不需要后缀
    viper.AddConfigPath("./")
    viper.AddConfigPath("/etc/viper_demo2/")
    //viper.SetConfigType("toml")

    viper.SetDefault("loggers", map[string]string{"path": "/var/log/viper_demo/access.log", "level": "debug"})
    viper.SetDefault("", "")

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal(err)
    }
    if err := viper.Unmarshal(&config); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", config)
}
