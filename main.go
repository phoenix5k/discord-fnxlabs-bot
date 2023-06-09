package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var token string = "token"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	bot := NewDiscoBot(token)
	if err := bot.Open(ctx); err != nil {
		log.Fatalln(err)
	}
	defer bot.Close()

	go func() {
		if err := bot.RunPlayer(ctx); err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	cancel()
}
