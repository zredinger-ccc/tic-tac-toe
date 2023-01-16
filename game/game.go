package game

import (
	"fmt"
	"log"
	"strconv"
)

type Game struct {
	PlacementCount int
	Turn           int
	Board          [3][3]string
}

func New() *Game {
	return &Game{
		PlacementCount: 0,
		Turn:           1,
		Board:          [3][3]string{},
	}
}

func (g *Game) PlayMove(loc []string) {
	x, err := strconv.Atoi(loc[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(loc[1])
	if err != nil {
		log.Fatal(err)
	}

	if g.Board[x-1][y-1] == "" {
		if g.Turn == 1 {
			g.Board[x-1][y-1] = "O"
		} else {
			g.Board[x-1][y-1] = "X"
		}
		g.Turn = (g.Turn + 1) % 2
		g.PlacementCount++
	} else {
		fmt.Println("already taken")
	}
}

func (g *Game) Render() {
	fmt.Println()
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[0]); j++ {
			if g.Board[i][j] == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(g.Board[i][j])
			}
		}
		fmt.Println()
	}
}

func (g *Game) IsGameOver() bool {
	if len(g.Board)*len(g.Board[0]) != g.PlacementCount {
		return false
	}

	return true
}

func (g *Game) IsGameWon() bool {
	// first check horizontal
	p1Count := 0
	p2Count := 0
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[i]); j++ {
			if g.Board[i][j] == "X" {
				p2Count++
			} else if g.Board[i][j] == "O" {
				p1Count++
			}
		}
		if p1Count == 3 {
			return true
		} else if p2Count == 3 {
			return true
		}
		p1Count = 0
		p2Count = 0
	}

	// next check vertical
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[i]); j++ {
			if g.Board[j][i] == "X" {
				p2Count++
			} else if g.Board[j][i] == "O" {
				p1Count++
			}
		}
		if p1Count == 3 {
			return true
		} else if p2Count == 3 {
			return true
		}
		p1Count = 0
		p2Count = 0
	}

	return false
	// then check diagonal
}
