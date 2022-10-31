package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

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
			fmt.Println("Uhmm command failed")
			// Handle the error
		}
	}
}
