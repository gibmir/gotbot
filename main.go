package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gibmir/gotbot/command"
	"github.com/gibmir/gotbot/config"
	"github.com/gibmir/gotbot/handler"
)

var (
	configuration *config.Config
)

func init() {
	reader := config.CreateReader()
	factory, err := config.CreateFactory(reader)
	if err != nil {
		log.Fatal("exception occurred while resolving config factory,", err)
	}
	c, err := factory.Create()
	if err != nil {
		log.Fatal("exception occurred while creating config,", err)
	}
	configuration = c
}

func main() {
	discord, err := discordgo.New("Bot " + configuration.Token)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
	}
	registry := command.NewRegistry()
	processor := command.NewDynamicProcessor(&registry)
	h := handler.NewDiscordHandler(&processor)
	discord.AddHandler(h.OnMessage)

	// In this example, we only care about receiving message events.
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("gotbot is ready to go :) Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
