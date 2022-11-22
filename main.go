package main

import (
	"fmt"
)

func main() {
	fmt.Println("----Football Scoreboard----")

	gamescore1 := newScoreBoard("Germany", "Italy")
	gamescore2 := newScoreBoard("France", "Portugal")
	gamescore3 := newScoreBoard("Spain", "Brazil")


	fmt.Println(gamescore1)
	fmt.Println(gamescore2)
	fmt.Println(gamescore3)

	fmt.Println("******Game Started******")

	gamescore1.updateScore(2,5)
	gamescore2.updateScore(4,4)
	gamescore3.updateScore(3,10)

	allScores := []scoreboard{gamescore1, gamescore2, gamescore3}
	fmt.Println(allScores)
  
	fmt.Println("******Game Summary******")

	for i := 0; i < len(allScores); i++ {
    if allScores[i].away_team_score == allScores[i].home_team_score{
      fmt.Printf("%s %d - %s %d\n", allScores[i].home_team, allScores[i].home_team_score, allScores[i].away_team, allScores[i].away_team_score)
    }
  }

  for i := 0; i < len(allScores); i++ {
    if allScores[i].away_team_score != allScores[i].home_team_score{
      fmt.Printf("%s %d - %s %d\n", allScores[i].home_team, allScores[i].home_team_score, allScores[i].away_team, allScores[i].away_team_score)
    }
  }
}