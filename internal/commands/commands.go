package commands

import "github.com/bwmarrin/discordgo"

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping pong",
		},
		{
			Name:        "modlog",
			Description: "Create a modlog thread",
		},
	}

	handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping":   PingHandler,
		"modlog": ModLog,
	}
)
