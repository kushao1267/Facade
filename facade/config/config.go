package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

type Redis struct {
	Addr        string
	Password    string
	DB          int
	PoolSize    int
	PoolTimeout int
}

type DefaultItem struct {
	Title string
	Image string
	ForceDefault bool
}

type Config struct {
	Title     string
	Expire    time.Duration
	Redis     map[string]Redis `toml:"redis"`
	ReturnMap map[string]DefaultItem
}

// AllConf 全局配置
var AllConf = &Config{}

func init() {
	if _, err := toml.DecodeFile("./config.toml", AllConf); err != nil {
		log.Fatal("加载配置失败", err)
	}
}
