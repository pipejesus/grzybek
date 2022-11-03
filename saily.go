package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/pipejesus/grzybek/sailygame"
)

var game *sailygame.SailyGame

func sailyCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	if data.Name != "saily" {
		return
	}

	err := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Saily Game here :)",
			},
		},
	)

	if err != nil {
		fmt.Println("Uhmm command failed for saily")
		return
	}

	options := data.Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))

	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if game == nil || game.HasEnded() {
		s.ChannelMessageSend(i.ChannelID, "Creating new game")
		s.ChannelMessage(i.ChannelID, "type `/saily` to play")
		game = sailygame.New()
		s.ChannelMessageSend(i.ChannelID, game.DisplayLevel())
	}

	// fmt.Println(i.Member)
	// fmt.Println(i.User)

	isNew := game.AddPlayer(i.Member)

	if isNew {
		s.ChannelMessageSend(i.ChannelID, "Player `"+i.Member.User.Username+"` joins Saily")
	}

	x, okx := optionMap["x"]
	y, oky := optionMap["y"]

	var msg string

	if okx && oky {
		_, msg = game.Shot(i.Member, int(x.IntValue()), int(y.IntValue()))
		s.ChannelMessageSend(i.ChannelID, game.DisplayLevel())
		s.ChannelMessageSend(i.ChannelID, msg)
	}

	s.ChannelMessageSend(i.ChannelID, game.GetPointsString())
}
