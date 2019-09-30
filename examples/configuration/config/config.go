package config

import (
    "encoding/json"
    "github.com/aaronize/grandonion/examples/configuration/config/logger"
    "github.com/aaronize/grandonion/examples/configuration/config/server"
    "log"
)

type Config struct {
    Service         string          `json:"service"`
    RunMode         string          `json:"run_mode"`

    ServerConfig    *server.Config    `json:"server"`
    LoggerConfig    *logger.Config    `json:"logger"`
}

func NewConfig() *Config {
    return &Config{
        Service: "grandonion",
        RunMode: "debug",

        ServerConfig: server.NewServerConfig(),
        LoggerConfig: logger.NewLoggerConfig(),
    }
}

func NewEmptyConfig() *Config {
    return &Config{}
}

func ParseConfig(path string) *Config {
    log.Printf("Using config.toml file: [%s]\n", path)

    config := NewConfig()

    if err := parseConfig(path, config); err != nil {
        log.Fatal(err)
    }

    return config
}

/*
从文件解析的时候，会覆盖默认值
即使文件中是零值也会解析并覆盖默认值
 */
func parseConfig(path string, cf *Config) error {
    configString := `
{
    "service": "oniond",
    "run_mode": "debug",
    "server": {
        "bind_address": "0.0.0.0"
    },
    "logger": {
        "file_path": "./log/access.log",
        "backups": 0,
        "level": "info"
    }
}
    `
    return json.Unmarshal([]byte(configString), cf)
}
