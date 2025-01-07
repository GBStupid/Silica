package main

import (
  "silica/internal/client"
  "silica/internal/config"

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

  client.GetSession().Open()
}
