package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"foodApp/cmd/customer/app"
)

const defaultConfPath = "./local.yaml"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := app.Application{}
	application.Init(ctx, defaultConfPath)

	application.Start(ctx)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sigterm
	application.Stop(ctx)
	defer func(cancel context.CancelFunc) {
		cancel()
		os.Exit(0)
	}(cancel)
}
