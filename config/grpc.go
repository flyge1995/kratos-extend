package config

import "time"

type GRPC struct {
	Network string        `mapstructure:"network"`
	Addr    string        `mapstructure:"addr"`
	Timeout time.Duration `mapstructure:"timeout"`
}
