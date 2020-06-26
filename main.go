package main

import (
	"log"
	"os"
	"time"

	"github.com/marcelo-lozoya/reddit-downloader/cli"
)

func main() {
	time.Now().Nanosecond()

	start := time.Now()

	li := cli.Cli{Commands: []cli.Command{}}

	li.Init(os.Args)

	log.Println("took ", time.Since(start))
}
