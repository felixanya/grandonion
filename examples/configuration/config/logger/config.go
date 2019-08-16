package logger

const (
    Info = "info"
    Debug = "debug"
    Error = "error"

    defaultFilePath = "/var/log/grandonion/server.log"
    defaultBackup = 10
    defaultLevel = Debug
)

var config *Config

type Config struct {
    FilePath    string      `json:"file_path"`
    Backups     int         `json:"backups"`
    Level       string      `json:"level"`
}

func NewLoggerConfig() *Config {
    config =  &Config{
        FilePath: defaultFilePath,
        Backups: defaultBackup,
        Level: defaultLevel,
    }

    return config
}

func GetLoggerConfig() *Config {
    return config
}
