package server

type Config struct {
    BindAddress     string      `json:"bind_address"`
    Port            int         `json:"port"`
}

func NewServerConfig() *Config {
    return &Config{
        BindAddress: "127.0.0.1",
        Port: 15100,
    }
}
