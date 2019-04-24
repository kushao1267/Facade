package config

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

type Redis struct {
	Addr        string
	Password    string
	DB          int
	PoolSize    int
	PoolTimeout int

	Expire time.Duration
}

type Config struct {
	Title     string
	Redis     map[string]Redis `toml:"redis"`
	ReturnMap map[string]DefaultItem
}

type DefaultItem struct {
	Title        string
	Image        string
	ForceDefault bool
}

// AllConf 全局配置
var AllConf = &Config{}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

func init() {
	AllConf.Read()
}
