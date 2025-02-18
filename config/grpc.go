package config

import "time"

type GRPC struct {
	Network string
	Addr    string
	Timeout time.Duration
}
