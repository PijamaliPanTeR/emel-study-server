package pkg

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type AppConfig struct {
	Server ServerConfig `mapstructure:"server"`
}

func LoadConfig(configPath string) (*AppConfig, error) {
	if configPath == "" {
		configPath = "conf/local.yml"
	}
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &cfg, nil
}

func MustGetwd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
