//go:build wireinject
// +build wireinject

package zap

import "github.com/google/wire"

var (
	LoggerProviderSet = wire.NewSet(NewLumberjack, NewZap, NewLogger, NewKV)
)
