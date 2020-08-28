package main

import (
	"log"
	"time"
)

func main() {
	appContainer := InitializeApp() // Initialize dependency injection container first!
	downloader := appContainer.downloader
	cli := appContainer.cli

	start := time.Now()

	downloader.MakeRequestForReddit(cli.GetArguments())

	log.Println("took ", time.Since(start))
}
