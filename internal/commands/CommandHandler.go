package commands

import (
	"github.com/Musik-Bot/Musik-Bot/config"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Command struct {
	Command     string                                                 `json:"command"`
	Description string                                                 `json:"description"`
	Executor    func(s *discordgo.Session, m *discordgo.MessageCreate) `json:"executor"`
	Permission  string                                                 `json:"permission"`
}

func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cfg, _ := config.Parse()
	if strings.HasPrefix(m.Content, ";;") {
		if m.ChannelID == cfg.ChannelID {
			command := strings.Split(m.Content, ";;")[1]
			AllCommands := GetAllCommands()
			if val, ok := AllCommands[command]; ok {
				val.Executor(s, m)
			} else {
				CommandDoesNotExist(s, m)
			}
		}
	}
}

func GetAllCommands() map[string]Command {
	m := make(map[string]Command)
	m["info"] = Command{"info", "The main info command", InfoCommand, "None"}
	m["help"] = Command{"help", "Help Command", HelpCommand, "None"}
	return m
}
