package sailygame

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type player struct {
	points int
	member *discordgo.Member
}

type SailyGame struct {
	players map[string]*player
	ended   bool
}

func New() *SailyGame {
	return &SailyGame{
		make(map[string]*player),
		false,
	}
}

func (g *SailyGame) HasEnded() bool {
	return g.ended
}

func (g *SailyGame) AddPlayer(member *discordgo.Member) {

	if p, exists := g.players[member.User.ID]; exists {
		p.points += 1

		return
	}

	g.players[member.User.ID] = &player{
		0,
		member,
	}
}

func (g *SailyGame) GetPointsString() (points string) {
	points = ""

	for _, player := range g.players {
		points += fmt.Sprintf("%s : %d points ", player.member.User.Username, player.points)
	}

	return
}
