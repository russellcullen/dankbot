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
	if m.Content == "KappaArjun" {
		replaceWithFile(s, m, "/Users/russell/Downloads/arjunkappa-small.png", "png")
	}
	if m.Content == "SombraDance" {
		replaceWithFile(s, m, "/Users/russell/Downloads/sombra-anim.gif", "gif")
	}
}

func replaceWithFile(s *discordgo.Session, m *discordgo.MessageCreate, filepath string, filetype string) {
	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		fmt.Println("Error deleting message: ", err)
	}
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening image: ", err)
	}
	_, err = s.ChannelFileSend(m.ChannelID, filetype, file)
	if err != nil {
		fmt.Println("Error sending spicy meme: ", err)
	}
}
