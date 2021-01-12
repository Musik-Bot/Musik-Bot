package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

func CommandDoesNotExist(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := embed.NewEmbed()
	embed.SetTitle("Invalid command")
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
	if err != nil {
		fmt.Println("Cannot send embed message to channel")
	}
}
