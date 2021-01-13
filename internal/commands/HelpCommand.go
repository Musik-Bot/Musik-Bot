package commands

import (
	"fmt"
	"github.com/Musik-Bot/Musik-Bot/config"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	commandList := GetAllCommands()
	emb := embed.NewEmbed()
	emb.Title = "Help"
	emb.Color = 0xADE81C
	emb.Description = "Help command"
	cfg, _ := config.Parse()
	channel, _ := s.Channel(cfg.CommandChannel)
	emb.AddField("Prefix:", cfg.Prefix)
	emb.AddField("Command channel:", channel.Name)
	for _, v := range commandList {
		emb.AddField(v.Command, "description: "+v.Description+"\npermission: "+v.Permission)
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	if err != nil {
		fmt.Println("Cannot execute help command")
	}
}
