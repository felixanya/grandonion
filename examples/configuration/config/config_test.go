package config

import (
    "fmt"
    "github.com/aaronize/grandonion/examples/configuration/config/logger"
    "testing"
)

func TestParseConfig(t *testing.T) {
    config := ParseConfig("/etc/grandonion/server.conf")

    fmt.Printf("\n---> %+v\n", config)
    fmt.Printf("---> %+v\n", config.LoggerConfig)
    fmt.Printf("---> %+v\n", config.ServerConfig)

    fmt.Println()
    fmt.Printf("---> %+v\n", logger.GetLoggerConfig())
}
