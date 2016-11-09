package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"google.golang.org/appengine"
	"os"
)

func main() {
	discord, err := discordgo.New("Bot MjQ1NTc3Nzk3MjQ0OTQ0Mzg1.CwORrg.9RU0tgVnRE3s41Y-W8Z0PU339q8")
	if err != nil {
		fmt.Println("Error authenticating: ", err)
	}
	discord.AddHandler(newMessage)
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("DankBot is now running.")
	appengine.Main()
}

func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content)
	switch m.Content {
	case "SombraDance":
		sendMessage(s, m.ChannelID, "http://i.imgur.com/lq3TwJi.gif")
	}
}

func sendMessage(s *discordgo.Session, channelId string, content string) {
	_, err := s.ChannelMessageSend(channelId, content)
	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}
