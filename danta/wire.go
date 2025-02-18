//go:build wireinject
// +build wireinject

package config

import "github.com/google/wire"

func ProvideContext() wire.ProviderSet {
	return wire.NewSet(NewContext, NewContextSide, NewCancelFuncSide)
}
