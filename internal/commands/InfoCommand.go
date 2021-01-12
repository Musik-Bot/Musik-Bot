package commands

import (
	"fmt"
	"github.com/Musik-Bot/Musik-Bot/internal/system"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
	"strconv"
)

func InfoCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	emb := embed.NewEmbed()
	emb.Title = "Information"
	emb.Description = "Some information about the bot."
	emb.Color = 0x1C76E8
	sysInfo := system.GetSystemInformation()
	emb.AddField("CPU:", sysInfo.CPU)
	emb.AddField("RAM:", strconv.Itoa(int(sysInfo.RAM_USED))+"/"+strconv.Itoa(int(sysInfo.RAM_All))+"MB")
	emb.AddField("SWAP::", strconv.Itoa(int(sysInfo.SWAP_USED))+"/"+strconv.Itoa(int(sysInfo.SWAP_All))+"MB")
	emb.AddField("disk:", strconv.Itoa(int(sysInfo.Disk))+"GB")
	emb.AddField("operating system:", sysInfo.Platform)
	emb.AddField("Bot version:", "v0.0.1-dev")
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	if err != nil {
		fmt.Println("Cannot send InformationEmbed")
	}
}
