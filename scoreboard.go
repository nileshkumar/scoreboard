package main

type scoreboard struct {
	home_team string
	away_team string
	home_team_score int8
	away_team_score int8
}

func newScoreBoard(home_team string, away_team string) scoreboard {
  score := scoreboard{
    home_team: home_team,
    away_team: away_team,
    home_team_score: 0,
    away_team_score: 0,
  }

  return score
}