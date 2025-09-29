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
	fmt.Println("LE BUT EST DE ALIGNEE 4 JETONS ")
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
func (g *Game) afficherTableau() {
	fmt.Println("\nJoueur 1 : X")
	fmt.Println("Joueur 2 : O")
	fmt.Println("Joueur courant :", g.currentPlayer)
	fmt.Println(" 0 1 2 3 4 5 6")
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print("|", g.grille[i][j])
		}
		fmt.Println("|")
	}
	fmt.Println("---------------")
}

func (g *Game) AddJeton(colonne int) bool {
	if colonne < 0 || colonne >= 7 {
		return false
	}

	for i := 5; i >= 0; i-- {
		if g.grille[i][colonne] == " " {
			g.grille[i][colonne] = g.currentPlayer
			g.nbrtour++
			return true
		}
	}
	return false
}

// FT3	= "fonction ajout de la fonction ahout des jetons"

func verifierVictoire(g Game) bool {
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			if g.grille[i][j] != " " &&
				g.grille[i][j] == g.grille[i][j+1] &&
				g.grille[i][j] == g.grille[i][j+2] &&
				g.grille[i][j] == g.grille[i][j+3] {
				return true
			}
		}
	}
	for j := 0; j < 7; j++ {
		for i := 0; i < 3; i++ {
			if g.grille[i][j] != " " &&
				g.grille[i][j] == g.grille[i+1][j] &&
				g.grille[i][j] == g.grille[i+2][j] &&
				g.grille[i][j] == g.grille[i+3][j] {
				return true
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if g.grille[i][j] != " " &&
				g.grille[i][j] == g.grille[i+1][j+1] &&
				g.grille[i][j] == g.grille[i+2][j+2] &&
				g.grille[i][j] == g.grille[i+3][j+3] {
				return true
			}
		}
	}
	for i := 3; i < 6; i++ {
		for j := 0; j < 4; j++ {
			if g.grille[i][j] != " " &&
				g.grille[i][j] == g.grille[i-1][j+1] &&
				g.grille[i][j] == g.grille[i-2][j+2] &&
				g.grille[i][j] == g.grille[i-3][j+3] {
				return true
			}
		}
	}

	return false
}
func (g *Game) verifierMatchNul() {
	if g.nbrtour >= 42 {
		fmt.Println("Match nul ! La grille est pleine.")
		g.reset()
	}
}
