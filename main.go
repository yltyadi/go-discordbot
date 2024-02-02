package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// discord bot token is stored as an env variablee
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		fmt.Println("Discord Bot token not found. Please set the DISCORD_TOKEN environment variable.")
		return
	}

	// creating a new discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// adding message handlers
	dg.AddHandler(messageCreate)

	// connecting to discord session
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
		return
	}

	fmt.Println("Bot is now running and listening to your commands. Press Ctrl+C to exit.")

	// part of the code that stops the bot when needed
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// closing the discord session before stopping the bot
	dg.Close()
}

// This function will be called whenever a new message is created
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// necessary so that the bot ignores the messages sent by itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// respond to "ping" with "pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong!")
	}
}
