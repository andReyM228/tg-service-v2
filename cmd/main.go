package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tg-service-v2/internal/app"

	_ "github.com/lib/pq"
)

const serviceName = "tg-service"

func main() {
	a := app.New(serviceName)
	a.Run(gracefulShutDown())
}

func gracefulShutDown() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal)

	signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, os.Interrupt)
	go func() {
		<-c
		log.Println("services stopped by gracefulShutDown")
		cancel()

	}()

	return ctx
}
