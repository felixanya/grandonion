package proxy

type Config struct {
    Addr        string   `json:"addr"`
    IsAnonymous bool     `json:"is_anonymous"`
    Debug       bool     `json:"debug"`
}

func NewConfig() *Config {
    return &Config{
        Addr: "127.0.0.1:15103",
        IsAnonymous: false,
        Debug: true,
    }
}
