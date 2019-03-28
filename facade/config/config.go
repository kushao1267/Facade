package config

import (
	"github.com/BurntSushi/toml"
)

type Redis struct {
	Addr        string
	Password    string
	DB          int
	PoolSize    int
	PoolTimeout int
}

type Config struct {
	Title          string
	Redis map[string]Redis `toml:"redis"`
}

// AllConf 全局配置
var AllConf Config

func init() {
	if _, err := toml.DecodeFile("config.toml", &AllConf); err != nil {
		panic("加载配置失败")
	}
}