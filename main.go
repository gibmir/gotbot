package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gibmir/gotbot/config"
	"github.com/gibmir/gotbot/handler"
)

var (
	Configuration config.Config
)

func init() {
	Configuration = config.Create()
}

func main() {
	discord, err := discordgo.New("Bot " + Configuration.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	discord.AddHandler(handler.OnMessage)

	// In this example, we only care about receiving message events.
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
