package main

import (
	"fmt"
)

func main() {
	fmt.Println("----Football Scoreboard----")
	fmt.Println("******Game Started******")

	gamescore1 := newScoreBoard("Germany", "Italy")
	gamescore2 := newScoreBoard("France", "Portugal")
	gamescore3 := newScoreBoard("Spain", "Brazil")

	fmt.Println(gamescore1)
	fmt.Println(gamescore2)
	fmt.Println(gamescore3)
  
}