package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"wanderlust/commands"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config := ParseConfig()

	discord, err := discordgo.New("Bot " + config.Discord.SecurityToken)
	if err != nil {
		log.Fatal(err)
	}

	// Add event handler
	discord.AddHandler(newMessage)

	// Open websocket
	discord.Open()

	// Run until process is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	//close websocket
	defer discord.Close()
}

func newMessage(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messages
	if message.Author.ID == session.State.User.ID {
		return
	}

	// handle commands
	switch {
	case message.Content == "~wanderlust", message.Content == "~help":
		commands.HandleHelp(session, message)
	}
}