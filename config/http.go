package config

import "time"

type HTTP struct {
	Network string        `mapstructure:"network"`
	Addr    string        `mapstructure:"addr"`
	Timeout time.Duration `mapstructure:"timeout"`
}
