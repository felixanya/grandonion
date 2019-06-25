package main

import (
	"sync"
)

type Config struct {
	Http 	*Http
	Redis 	*Redis
	DB 		*DB
	Logger 	*Logger
}

type Http struct {
	Listen 		string
}

type Redis struct {
	Host 	string
	Qname 	string

}

type DB struct {
	Host 	string
	Port 	int
	User 	string
	Passwd 	string
	Database string
}

type Logger struct {
	Path 	string
}

const DEFAULT_CONF_PATH = "../config/conf.toml"

var (
	cfg 	*Config
	once 	sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {

	})

	return cfg
}
