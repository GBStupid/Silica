package commands

import (
  "silica/internal/types"
  "github.com/bwmarrin/discordgo"
)

func RegisterSlash(b types.Bot) {
  session := b.GetSession()
  session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
    ModLog(b, s, i) 
  })
}
