package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Bot *Bot
}

type Bot struct {
	Token       string `yaml:"token"`
	EnableProxy bool   `yaml:"enableProxy"`
	Proxy       string `yaml:"proxy"`
}

func InitConfig() *Config {
	AppConfig := &Config{}
	content, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}

	fmt.Println(string(content))
	if yaml.Unmarshal(content, &AppConfig) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	return AppConfig
}
