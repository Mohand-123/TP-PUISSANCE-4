package main

import (
	"fmt"
)

type Game struct {
	player01      string
	player02      string
	grille        [6][7]string
	currentPlayer string
	nbrtour       int
}

func initGame() Game {
	fmt.Println("========== BIENVENUE DANS LE PUISSANCE 4 ===========")
	fmt.Println("LE JOUEUR 1 AURA LES X, LE JOUEUR 2 AURA LES O")
	fmt.Println("BONNE CHANCE JOUEUR !")
	var grille [6][7]string
	for i := range grille {
		for j := range grille[i] {
			grille[i][j] = " "
		}
	}

	return Game{
		player01:      "X",
		player02:      "O",
		grille:        grille,
		currentPlayer: "X",
		nbrtour:       0,
	}
}
