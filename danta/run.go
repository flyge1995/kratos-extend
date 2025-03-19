package danta

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
)

type AppBuild func() (*kratos.App, func(), error)

func Run(build AppBuild) error {
	flag.Parse()

	app, cleanup, err := build()
	if err != nil {
		return err
	}
	defer cleanup()

	return app.Run()
}
