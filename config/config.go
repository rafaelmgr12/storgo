package config

import (
	"github.com/spf13/viper"
)

type Ports struct {
	Server1 int `mapstructure:"server1"`
	Server2 int `mapstructure:"server2"`
	Server3 int `mapstructure:"server3"`
}

type Config struct {
	AppName string `mapstructure:"app_name"`
	Port    Ports  `mapstructure:"port"`
}

func LoadConfig(path string) *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	viper.SetDefault("app_name", "Storgo")
	viper.SetDefault("port.server1", 3000)
	viper.SetDefault("port.server2", 7000)
	viper.SetDefault("port.server3", 5000)

	if err := viper.ReadInConfig(); err != nil {
		panic("error in reading config file: " + err.Error())
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		panic("unable to unmarshal config into struct: " + err.Error())
	}

	return &cfg
}
