package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"foodApp/cmd/restaurant/app"
)

const (
	defaultConfPath = "./local.yaml"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", defaultConfPath, "config file to load")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := app.Application{}
	application.Init(ctx, configFile)

	application.Start(ctx)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sigterm
	defer func(cancel context.CancelFunc) {
		cancel()
		os.Exit(0)
	}(cancel)
}
