package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token   string
	AppId   string
	GuildId string
	GifHost string
	OutUrl  string = "http://srv12.mikr.us:30472"
)

func init() {
	flag.StringVar(&Token, "t", "NzQxMDY3NzYzODQ5NjI1NjUx.GG8h-D.ZFbbVO0hIjBSG3Q-PbZStOG8vbWhmoyIi2D4cc", "Bot Token")
	flag.StringVar(&GifHost, "g", "http://localhost:30472", "The URL of the GIF app")
	flag.StringVar(&AppId, "a", "", "Discord APP ID")
	flag.StringVar(&GuildId, "s", "", "Discord Guild ID (Server ID)")
	flag.Parse()
}

func main() {

	dg, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// create commands
	_, err = dg.ApplicationCommandBulkOverwrite(AppId, GuildId, []*discordgo.ApplicationCommand{
		{
			Name:        "saily",
			Description: "Creates the saily game",
		},
		{
			Name:        "say",
			Description: "Outputs animated paszczak GIF",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "git-user-name",
					Description: "Name of the GitHub user account",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "message",
					Description: "Message to put in the bubble",
					Required:    true,
				},
			},
		},
	})

	// Handle the error of creating commands
	if err != nil {
		fmt.Println("Unable to create server commands")
		return
	}

	dg.AddHandler(sailyCommands)
	dg.AddHandler(sayCommand)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	if err := dg.Open(); err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
