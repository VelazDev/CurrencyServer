package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"app"`
}

func LoadConfig() (*Config, error) {
	conf := viper.New()

	config_name := fmt.Sprintf("%s.yaml", getWebEnv())
	fmt.Println(config_name)
	conf.SetConfigFile(config_name)
	conf.SetConfigType("yaml")
	conf.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	conf.AutomaticEnv()

	if err := conf.ReadInConfig(); err != nil {
		return nil, err
	}

	c := &Config{}
	err := conf.Unmarshal(c)

	log.Println(c)

	return c, err
}

const defaultEnv = "dev"

func getWebEnv() string {
	env := os.Getenv("WEB_ENV")
	if env != "" {
		return env
	}

	return defaultEnv
}
