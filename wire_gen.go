// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/marcelo-lozoya/reddit-downloader/cli"
	"github.com/marcelo-lozoya/reddit-downloader/downloader"
)

// Injectors from wire.go:

func InitializeApp() *appContainer {
	iCli := cli.InitCli()
	iDirectoryHelper := downloader.InitDirectoryHelper()
	iDownloader := downloader.InitDownloader(iDirectoryHelper)
	mainAppContainer := &appContainer{
		cli:        iCli,
		downloader: iDownloader,
	}
	return mainAppContainer
}

// wire.go:

var containerSet = wire.NewSet(wire.Struct(new(appContainer), "*"), cli.InitCli, downloader.InitDownloader, downloader.InitDirectoryHelper)

type appContainer struct {
	// put any dependencies here that you want the main.go file to use
	cli        cli.ICli
	downloader downloader.IDownloader
}
