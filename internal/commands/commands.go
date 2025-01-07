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
      Options: []*discordgo.ApplicationCommandOption{
        {
          Name:        "user",
          Description: "user to modlog",
          Type:        discordgo.ApplicationCommandOptionMentionable,
          Required:    true,
        },
        {
          Name:        "action",
          Description: "action for modlog",
          Type:        discordgo.ApplicationCommandOptionString,
          Required:    true,
        },
        {
          Name:        "reason",
          Description: "reason for modlog",
          Type:        discordgo.ApplicationCommandOptionString,
          Required:    true,
        },
        {
          Name:        "proof",
          Description: "proof for the modlog",
          Type:        discordgo.ApplicationCommandOptionAttachment,
          Required:    false,
        },
      },
		},
	}

	handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping":   PingHandler,
		"modlog": ModLog,
	}
)
