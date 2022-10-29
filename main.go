package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token   string
	GifHost string
	OutUrl  string = "http://srv12.mikr.us:30472"
)

func init() {

	flag.StringVar(&Token, "t", "NzQxMDY3NzYzODQ5NjI1NjUx.GG8h-D.ZFbbVO0hIjBSG3Q-PbZStOG8vbWhmoyIi2D4cc", "Bot Token")
	flag.StringVar(&GifHost, "g", "http://localhost:30472", "The URL of the GIF app")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.

	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "say#") {
		return
	}

	msg := strings.Split(strings.ReplaceAll(m.Content, "\n", " "), "say#")

	fmt.Println(msg)
	fmt.Println(len(msg))

	if len(msg) != 2 {
		return
	}

	bas := GifHost + "/czo?gituser=greg-develtio&msg="
	bas_out := OutUrl + "/czo?gituser=greg-develtio&msg="
	gif_msg := url.QueryEscape(msg[1])

	response, err := http.Get(bas + gif_msg)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return
	}

	_, err = s.ChannelFileSend(m.ChannelID, "grzybotsays.gif", response.Body)

	if err != nil {
		return
	}
	// s.ChannelMessageSend(m.ChannelID, bas+gif_msg)
	s.ChannelMessageSend(m.ChannelID, "`"+bas_out+gif_msg+"`")
}
