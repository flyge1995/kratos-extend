package config

import "time"

type HTTP struct {
	Network string
	Addr    string
	Timeout time.Duration
}
