package commands

import (
	"log"
	"silica/internal/types"

	"github.com/bwmarrin/discordgo"
)

func RegisterSlash(b types.Client) {
	s := b.GetSession()
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handler, ok := handlers[i.ApplicationCommandData().Name]
		if !ok {
			return
		}
		handler(s, i)
	})
}
