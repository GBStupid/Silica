package client

import (
	"flag"

	"github.com/bwmarrin/discordgo"

	"silica/internal/config"
)

var (
	GuildID = flag.String("guild", "1257611555198074922", "Guild ID")
)

type Bot struct {
	Session *discordgo.Session
	Config  *config.Config
}

func NewSession(config *config.Config) (*Bot, error) {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}
	return &Bot{session, config}, nil
}

func (b *Bot) GetConfig() *config.Config {
	return b.Config
}

func (b *Bot) GetSession() *discordgo.Session {
	return b.Session
}
