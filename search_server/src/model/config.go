package model

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

// Config 定义Config结构体
type Config struct {
	// kafka配置
	Kafka struct {
		Ip    []string `yml:"ip"`
		Topic string   `yml:"topic"`
	} `yml:"kafka"`

	// Nacos配置
	Nacos struct {
		Ip          string `yml:"ip"`
		Port        int    `yml:"port"`
		Path        string `yml:"path"`
		NameSpaceId string `yml:"timeoutMs"`
		TimeoutMs   uint64 `yml:"timeoutMs"`
		Group       string `yml:"group"`
		A           string `yml:"a"`
	} `yml:"nacos"`

	// Redis配置
	Redis struct {
		Ip          string        `yml:"ip"`
		Port        int           `yml:"port"`
		MaxIdle     int           `yml:"maxIdle"`
		MaxActive   int           `yml:"maxActive"`
		IdleTimeout time.Duration `yml:"idleTimeout"`
		Network     string        `yml:"network"`
	} `yml:"redis"`
}

func NewConfig() *Config {
	return &Config{}
}

// ReadConfig 读取配置映射到结构体
func (config *Config) ReadConfig() (*Config, error) {
	log.Println("开始加载配置文件")
	file, err := os.ReadFile("src/config/config.yml")
	if err != nil {
		log.Println("读取文件config.yml发生错误")
		return nil, err
	}
	if yaml.Unmarshal(file, config) != nil {
		log.Println("解析文件config.yml发生错误")
		return nil, err
	}
	log.Println("成功加载配置文件")
	return config, nil
}
