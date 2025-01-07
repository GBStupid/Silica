package types

import (
	"silica/internal/config"

	"github.com/bwmarrin/discordgo"
)

type Client interface {
	GetConfig() *config.Config
	GetSession() *discordgo.Session
}
