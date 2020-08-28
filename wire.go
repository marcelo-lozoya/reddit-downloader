// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/marcelo-lozoya/reddit-downloader/cli"
	"github.com/marcelo-lozoya/reddit-downloader/downloader"
)

func InitializeApp() *appContainer {
	panic(wire.Build(containerSet))
}

var containerSet = wire.NewSet(
	// provider for our app's DI container
	wire.Struct(new(appContainer), "*"),

	// providers for everything we want in the container
	cli.InitCli,
	downloader.InitDownloader,
	downloader.InitDirectoryHelper,
)

type appContainer struct {
	// put any dependencies here that you want the main.go file to use
	cli        cli.ICli
	downloader downloader.IDownloader
}
