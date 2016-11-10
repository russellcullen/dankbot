package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/coolbrow/dankbot/reddit"
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
	args := strings.Split(m.Content, " ")
	switch args[0] {
	case "/sombra":
		sendMessage(s, m.ChannelID, "http://i.imgur.com/lq3TwJi.gif")
	case "/reddit":
		var url string
		if len(args) >= 3 {
			url = reddit.RandomSearch(args[1], args[2])
		} else if len(args) == 2 {
			url = reddit.Random(args[1])
		}
		if url != "" {
			sendMessage(s, m.ChannelID, url)
		} else {
			sendMessage(s, m.ChannelID, "No results found")
		}
	}
}

func sendMessage(s *discordgo.Session, channelId string, content string) {
	_, err := s.ChannelMessageSend(channelId, content)
	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}
