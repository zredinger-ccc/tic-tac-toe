package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"tic-tac-toe/game"
)

func main() {
	g := game.New()

	for {
		r := bufio.NewReader(os.Stdin)
		turn, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		loc := strings.Split(strings.TrimRight(turn, "\r\n"), ",")
		g.PlayMove(loc)

		g.Render()
		if g.IsGameWon() {
			fmt.Println("GAME WON!")

			break
		} // a comment

		if g.IsGameOver() {
			fmt.Println("Game Over, no winner!!")
			break
		}
	}
}
