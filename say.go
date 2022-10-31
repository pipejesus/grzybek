// Handles the say functionality
// e.g. say#Your Very Long Message

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Handles receiving the command
// and responds with an interaction message
// and a GIF file send
func sayCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))

	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	var gitUserName string
	var message string

	if option, ok := optionMap["git-user-name"]; ok {
		gitUserName = option.StringValue()
	}

	if option, ok := optionMap["message"]; ok {
		message = option.StringValue()
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("One moment please!"),
		},
	})

	sendGifToChannel(s, i, gitUserName, message)
}

// Creates the GIF using the GIF REST API
// and sends the ready GIF to channel publicly
func sendGifToChannel(s *discordgo.Session, i *discordgo.InteractionCreate, gitHubUser string, message string) {

	msg := url.QueryEscape(strings.ReplaceAll(message, "\n", " "))

	gifferUrl := GifHost + fmt.Sprintf("/czo?gituser=%s&msg=%s", gitHubUser, msg)
	gifferUrlDebug := OutUrl + fmt.Sprintf("/czo?gituser=%s&msg=%s", gitHubUser, msg)

	response, err := http.Get(gifferUrl)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return
	}

	_, err = s.ChannelFileSend(i.ChannelID, "grzybotsays.gif", response.Body)

	if err != nil {
		return
	}

	s.ChannelMessageSend(i.ChannelID, "`"+gifferUrlDebug+"`")
}
