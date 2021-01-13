package main

import (
	"fmt"
	BotConf "github.com/Musik-Bot/Musik-Bot/config"
	"github.com/Musik-Bot/Musik-Bot/internal/commands"
	"github.com/Musik-Bot/Musik-Bot/internal/mysql"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mysql.InitDatabase()
	config, err := BotConf.Parse()
	if err != nil {
		fmt.Println("Error while parsing bot configuration")
	}
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("creating discord session via websocket failed", err)
		return
	}
	session.AddHandler(commands.CommandHandler)
	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	err = session.Open()
	if err != nil {
		fmt.Println("Cannot connect to discord websocket")
		return
	}

	fmt.Println("The bot is running now. Terminate by using CTRL-C")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	session.Close()
}
