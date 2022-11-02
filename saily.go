package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/pipejesus/grzybek/sailygame"
)

var game *sailygame.SailyGame

func sailyCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	switch data.Name {
	case "saily":
		err := s.InteractionRespond(
			i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hello world from the saily command!",
				},
			},
		)
		if err != nil {
			fmt.Println("Uhmm command failed for saily")
			return
		}

		if game == nil || game.HasEnded() {
			game = sailygame.New()
		}

		// fmt.Println(i.Member)
		// fmt.Println(i.User)

		game.AddPlayer(i.Member)
		s.ChannelMessageSend(i.ChannelID, "Player `"+i.Member.User.Username+"` joins Saily")
		s.ChannelMessageSend(i.ChannelID, game.GetPointsString())
	}
}
