package main

import (
	"./reddit"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
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
		var msg string
		if len(args) >= 3 {
			msg = reddit.RandomSearch(args[1], args[2])
		} else if len(args) == 2 {
			msg = reddit.Random(args[1])
		}
		if msg == "" {
			msg = "*No results found.*"
		}
		sendMessage(s, m.ChannelID, msg)
	case "/rip":
		name := strings.Join(args[1:], " ")
		rip := fmt.Sprintf("http://www.tombstonebuilder.com/generate.php?top1=RIP&top3=%s", name)
		sendMessage(s, m.ChannelID, rip)
	}
}

func sendMessage(s *discordgo.Session, channelID string, content string) {
	_, err := s.ChannelMessageSend(channelID, content)
	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}
