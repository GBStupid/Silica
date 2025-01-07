package types

import ( 
  "silica/internal/config"
  "github.com/bwmarrin/discordgo"
)

type Bot interface {
  GetConfig() *config.Config
  GetSession() *discordgo.Session
}
