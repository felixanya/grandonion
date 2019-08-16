package sequential

type Config struct {
    Service         string          `json:"service"`
    RunMode         string          `json:"run_mode"`

    ServerConfig    ServerConfig    `json:"server"`
    LoggerConfig    LoggerConfig    `json:"logger"`
}

type ServerConfig struct {
    BindAddress     string      `json:"bind_address"`
    Port            int         `json:"port"`
}

type LoggerConfig struct {
    FilePath    string      `json:"file_path"`
    Backups     int         `json:"backups"`
    Level       string      `json:"level"`
}


