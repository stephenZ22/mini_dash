package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port     int    `mapstructure:"port"`
	LogLevel string `mapstructure:"log_level"`
}

type DatabaseConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Database    string `mapstructure:"database"`
	MaxConnects int    `mapstructure:"max_connections"`
}

type MiniConfig struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

var Cfg MiniConfig

func LoadConfig(configPath string) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error loading config file: %w", err))
	}

	// 绑定到结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(fmt.Errorf("unable to decode config into struct: %w", err))
	}

	fmt.Printf("Configuration loaded successfully file path: %s\n", configPath)
}
