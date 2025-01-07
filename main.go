package main

import (
	"log"
	"os"
	"os/signal"
	"silica/internal/client"
	"silica/internal/commands"
	"silica/internal/config"
	"syscall"

	"github.com/bwmarrin/discordgo"
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

	client.GetSession().AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Silica is ready")
	})

	if err := client.GetSession().Open(); err != nil {
		panic(err)
	}

  log.Println("Deleting commands")
  commands.DeleteSlash(client)
	log.Println("Registering commands")
	go commands.RegisterSlash(client)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}
