package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ModLog(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channelID := i.ChannelID

	_, err := s.ThreadStart(channelID, "New Thread", discordgo.ChannelTypeGuildPublicThread, 60)
	if err != nil {
		log.Fatalf("Failed to create thread: %v", err)
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Thread created",
		},
	})
}
