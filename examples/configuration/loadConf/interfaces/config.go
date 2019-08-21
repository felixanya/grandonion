package configure

import (
    "github.com/BurntSushi/toml"
    "path/filepath"
)

// TODO 显然不能使用接口解析配置信息的，所以这里是错误的
type Config struct {
    Producer        Producer    `toml:"producer"`
}

func ParseConfig(path string) (*Config, error) {
    confPath, err := filepath.Abs(path)
    if err != nil {
        return nil, err
    }

    config := &Config{}
    _, err = toml.DecodeFile(confPath, config)
    if err != nil {
        return nil, err
    }

    return config, nil
}
