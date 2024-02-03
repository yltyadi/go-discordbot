package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	Token  = "MTIwMjg0NjU4Njk1NTIzOTQ3NA.GaF0wx.gL2IqafsNHJx7WYLVQAhcuvwdESmOX-39uC8qo"
	Prefix = "!"
)

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// Wait for a signal to exit
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the message starts with the command prefix
	if strings.HasPrefix(m.Content, Prefix) {
		fmt.Println("yep")
		// Parse the command
		command := strings.Split(strings.ToLower(m.Content)[1:], " ")

		switch command[0] {
		case "help":
			helpCommand(s, m)
		case "weather":
			weatherCommand(s, m, command[1:])
		case "translate":
			translateCommand(s, m, command[1:])
		default:
			s.ChannelMessageSend(m.ChannelID, "Unknown command. Type !help for a list of commands.")
		}
	}
}

func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpMessage := "Available commands:\n" +
		"!help - Display this help message.\n" +
		"!weather <location> - Get current weather information for the specified location.\n" +
		"!translate <language code> <text> - Translate the text to the specified language."

	s.ChannelMessageSend(m.ChannelID, helpMessage)
}

func weatherCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "Weather command...")
}

func translateCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "Translation command...")
}
