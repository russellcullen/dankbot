package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/coolbrow/dankbot/images"
	"github.com/coolbrow/dankbot/reddit"
	"github.com/coolbrow/dankbot/status"
	"google.golang.org/appengine"
	"net/url"
	"strings"
	"time"
)

const usage string = "\n" +
	"*DankBot Commands:*\n\n" +
	"`!help`\n" +
	"    Shows this text.\n\n" +
	"`!reddit [subreddit]`\n" +
	"    Returns random top result of given subreddit.\n\n" +
	"`!reddit [subreddit] [query]`\n" +
	"    Searches given subreddit for query and returns random top result.\n\n" +
	"`!rip [name]`\n" +
	"    RIP pic\n\n" +
	"`!retro [your], [text], [here]`\n" +
	"    Generates retro text\n\n" +
	"`!sombra`\n" +
	"    ~~Party~~ Fiesta time\n\n"

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

	statusUpdate := func(s string) {
		discord.UpdateStatus(0, s)
	}
	statusUpdate(status.RandomStatus())
	status.TickerStatus(time.Hour*6, statusUpdate)

	fmt.Println("DankBot is now running.")
	appengine.Main()
}

func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")
	switch args[0] {
	case "!sombra":
		sendMessage(s, m.ChannelID, images.Sombra())
	case "!reddit":
		var msg string
		if len(args) >= 3 {
			msg = reddit.RandomSearch(args[1], strings.Join(args[1:], " "))
		} else if len(args) == 2 {
			msg = reddit.Random(args[1])
		} else {
			msg = "*Must provide a subreddit.*"
		}
		if msg == "" {
			msg = "*No results found.*"
		}
		sendMessage(s, m.ChannelID, msg)
	case "!rip":
		name := url.QueryEscape(strings.Join(args[1:], " "))
		sendMessage(s, m.ChannelID, images.GenerateRIP(name))
	case "!retro":
		lines := strings.Split(strings.Join(args[1:], " "), ",")
		var text1, text2, text3 string
		switch {
		case len(lines) >= 3:
			text1 = lines[0]
			text2 = lines[1]
			text3 = strings.Join(lines[2:], " ")
		case len(lines) == 2:
			text1 = lines[0]
			text2 = lines[1]
		case len(lines) == 1:
			text2 = lines[0]
		}
		sendMessage(s, m.ChannelID, images.GenerateRetro(text1, text2, text3))
	case "!help":
		sendMessage(s, m.ChannelID, usage)
	}

}

func sendMessage(s *discordgo.Session, channelID string, content string) {
	_, err := s.ChannelMessageSend(channelID, content)
	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}
