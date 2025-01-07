package main

import (
	"log"
	"os"
	"os/signal"
	"silica/internal/client"
	"silica/internal/config"
	"syscall"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	client, err := client.NewSession(config)
	if err != nil {
		panic(err)
	}

	if err := client.GetSession().Open(); err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Silica is running")
	<-sc
}
