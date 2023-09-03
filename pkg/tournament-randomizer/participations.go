package main

type Players struct {
	Players []string
}

func NewPlayers(listOfPlayers []string) Players {
	return Players{Players: listOfPlayers}
}
