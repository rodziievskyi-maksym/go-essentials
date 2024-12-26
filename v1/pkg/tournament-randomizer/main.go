package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var players = []string{
	"Max",
	"Misha",
	"Oleja",
	//"Artem",
	//"Igor",
	//"Andrey",
	//"Nikita",
	//"Dima",
}

type Draft struct {
	Players
	GodTear
	SecondTier
}

func NewDraft(p Players, gt GodTear, st SecondTier) Draft {
	return Draft{
		Players:    p,
		GodTear:    gt,
		SecondTier: st,
	}
}

func main() {
	newPlayers := NewPlayers(players)
	godTier := NewGodTear(ReceiveGodTearTeams())
	secondTier := NewSecondTier(ReceiveSecondTierTeams())

	newDraft := NewDraft(newPlayers, godTier, secondTier)
	_ = newDraft.Draft()
}

func (d *Draft) Draft() *Set {
	newSet := NewSet()

	playersCount := len(d.Players.Players)

	for i := 0; i < playersCount; i++ {
		var godTeam string
		var secondTeam string
		//Randomly pick a number
		playerIndex := RandomFunc(len(d.Players.Players))
		pickedPlayer := d.Players.Players[playerIndex]

		//Pick a league
		randomGodLeagueIndex := RandomFunc(len(d.GodTear.Teams))
		godLeague := d.GodTear.Teams[randomGodLeagueIndex]
		if len(godLeague) == 0 {
			d.GodTear.Teams = RemoveLeague(d.GodTear.Teams, randomGodLeagueIndex)
			randomGodLeagueIndex = RandomFunc(len(d.GodTear.Teams))
			godLeague = d.GodTear.Teams[randomGodLeagueIndex]
		}

		randomGodLeagueTeamIndex := RandomFunc(len(godLeague))
		godTeam = godLeague[randomGodLeagueTeamIndex]

		//remove team
		d.GodTear.Teams[randomGodLeagueIndex] = Remove(godLeague, randomGodLeagueTeamIndex)

		newSet.Add(pickedPlayer, godTeam)

		//TODO make one method
		//Pick a league
		randomSecondLeagueIndex := RandomFunc(len(d.SecondTier.Teams))
		secondLeague := d.SecondTier.Teams[randomSecondLeagueIndex]
		if len(secondLeague) == 0 {
			d.SecondTier.Teams = RemoveLeague(d.SecondTier.Teams, randomSecondLeagueIndex)
			randomSecondLeagueIndex = RandomFunc(len(d.SecondTier.Teams))
			secondLeague = d.SecondTier.Teams[randomSecondLeagueIndex]
		}

		randomSecondLeagueTeamIndex := RandomFunc(len(secondLeague))
		secondTeam = secondLeague[randomSecondLeagueTeamIndex]

		//remove team
		d.SecondTier.Teams[randomSecondLeagueIndex] = Remove(secondLeague, randomSecondLeagueTeamIndex)

		newSet.Add(pickedPlayer, secondTeam)

		//Remove player from slice by index
		d.Players.Players = Remove(d.Players.Players, playerIndex)
	}

	//map data into readable format
	writeIntoFile(printSet(newSet), createEmptyFile())
	return newSet
}

// TODO: write generic instead of two methods
func Remove(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func RemoveLeague(s []Teams, index int) []Teams {
	return append(s[:index], s[index+1:]...)
}

func printSet(set *Set) string {
	var str string
	for key, value := range set.items {
		str += fmt.Sprintf("%s: plays for %v\n", key, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
	}

	return str
}

func RandomFunc(num int) int {
	source := rand.NewSource(time.Now().UnixNano())
	newRandom := rand.New(source)

	return newRandom.Intn(num)
}

type Set struct {
	items map[string][]string
}

func NewSet() *Set {
	return &Set{items: make(map[string][]string)}
}

func (s *Set) Has(player string) bool {
	_, ok := s.items[player]
	return ok
}

func (s *Set) Add(player string, team string) {
	s.items[player] = append(s.items[player], team)
}
