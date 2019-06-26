package up2redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Listen   string
	Logger   LoggerC
	Redis    RedisC
	PoolSize int 		`json:"pool_size"`
}

// load config
func (c *Config) LoadConfig(path string) error {
	if path == "" {
		return errors.New("path is null")
	}
	if !IsPathExist(path) {
		return errors.New(fmt.Sprintf("%s file not found", path))
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New(fmt.Sprintf("load %s file error. %s", path, err.Error()))
	}

	if len(content) == 0 {
		return errors.New(fmt.Sprintf("file %s is empty", path))
	}

	if err := json.Unmarshal(content, c); err != nil {
		return errors.New(fmt.Sprintf("parse %s file error. %s", path, err.Error()))
	}
	return nil
}

func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}


