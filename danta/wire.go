//go:build wireinject
// +build wireinject

package danta

import "github.com/google/wire"

var (
	ContextProviderSet = wire.NewSet(NewContext, NewContextSide, NewCancelFuncSide)
)
