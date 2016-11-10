package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/coolbrow/dankbot/haiku"
	"google.golang.org/appengine"
	"strings"
)

func main() {
	token := flag.String("token", "", "Client token")
	flag.Parse()

	if *token == "" {
		fmt.Println("Must provide client token. Run with -h for more info")
		return
	}
	discord, err := discordgo.New("Bot " + *token)
	if err != nil {
		fmt.Println("Error authenticating: ", err)
		return
	}
	discord.AddHandler(newMessage)
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	fmt.Println("DankBot is now running.")
	appengine.Main()
}

func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "SombraDance" {
		sendMessage(s, m.ChannelID, "http://i.imgur.com/lq3TwJi.gif")
	} else if strings.HasPrefix(m.Content, "/haiku") {
		query := strings.SplitN(m.Content, " ", 2)[1]
		url := haiku.TopUrl(query)
		if url != "" {
			sendMessage(s, m.ChannelID, url)
		} else {
			sendMessage(s, m.ChannelID, "No youtube haiku found")
		}
	}
}

func sendMessage(s *discordgo.Session, channelId string, content string) {
	_, err := s.ChannelMessageSend(channelId, content)
	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}
