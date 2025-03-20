package config

import (
	"github.com/spf13/viper"
)

type Option func(v *viper.Viper)

func NewConfig(o ...Option) *Config {
	v := viper.New()
	v.AddConfigPath("configs")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	for _, option := range o {
		option(v)
	}
	return &Config{v: v}
}

type Config struct {
	v *viper.Viper
}

func (c *Config) LoadEnvAndConfigFile(cfgFile string, cfgObj any) error {
	if cfgFile != "" {
		c.v.SetConfigFile(cfgFile)
	}

	c.v.AutomaticEnv()

	err := c.v.ReadInConfig()
	if err != nil {
		return err
	}

	err = c.v.Unmarshal(cfgObj)
	return err
}

func (c *Config) LoadConfigFile(cfgFile string, cfgObj any) error {
	if cfgFile != "" {
		c.v.SetConfigFile(cfgFile)
	}

	err := c.v.ReadInConfig()
	if err != nil {
		return err
	}

	err = c.v.Unmarshal(cfgObj)
	return err
}
