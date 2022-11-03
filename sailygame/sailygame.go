package sailygame

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

type player struct {
	points int
	member *discordgo.Member
}

type seafield [10][10]int

type SailyGame struct {
	players map[string]*player
	ended   bool
	sea     *seafield
}

func New() *SailyGame {
	return &SailyGame{
		make(map[string]*player),
		false,
		createSeafield(),
	}
}

func createSeafield() (sf *seafield) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	sf = &seafield{}
	for y := 0; y < len(sf); y++ {
		for x := 0; x < len(sf[0]); x++ {
			if rn := r1.Intn(100); rn > 50 {
				sf[y][x] = 1
			}
		}
	}
	return
}

func (g *SailyGame) HasEnded() bool {
	return g.ended
}

func (g *SailyGame) GetPlayerByMember(member *discordgo.Member) (*player, bool) {
	if player, exists := g.players[member.User.ID]; exists {
		return player, true
	}

	return nil, false
}

func (g *SailyGame) AddPlayer(member *discordgo.Member) bool {
	if _, exists := g.GetPlayerByMember(member); exists {
		return false
	}

	g.players[member.User.ID] = &player{
		0,
		member,
	}

	return true
}

func (g *SailyGame) AddPoint(member *discordgo.Member) {
	if player, exists := g.GetPlayerByMember(member); exists {
		player.points += 1
	}
}

func (g *SailyGame) DisplayLevel() string {
	str := ""
	for y := 0; y < len(g.sea); y++ {
		row := ""
		for x := 0; x < len(g.sea[0]); x++ {
			switch g.sea[y][x] {
			case 0:
				row += "[ ]"
			case 1:
				row += "[ ]"
			case 2:
				row += "[o]"
			}
		}
		str += row + "\n"
	}
	return "```" + str + "```"
}

func (g *SailyGame) Shot(member *discordgo.Member, x int, y int) (bool, string) {
	if x > 9 || x < 0 || y > 9 || y < 0 {
		return false, "Not on board!"
	}

	if g.sea[y][x] == 1 {
		g.AddPoint(member)
		g.sea[y][x] = 2
		return true, "Hit!"
	} else if g.sea[y][x] == 2 {
		return false, "Already discovered!"
	}

	return false, "Miss!"
}

func (g *SailyGame) GetPointsString() (points string) {
	points = ""

	for _, player := range g.players {
		points += fmt.Sprintf("%s : %d points ", player.member.User.Username, player.points)
	}

	return
}
