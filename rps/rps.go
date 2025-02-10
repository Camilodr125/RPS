package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMesaages = []string{
	"Well done!!!",
	"Good job :)",
	"You shound buy a lottery ticket",
}

var loseMesaages = []string{
	"Sorry :(",
	"Try Again...",
	"This is not your day",
}

var drawMesaages = []string{
	"It's a draw",
	"You're thinking the same",
	"Nobody wins",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)
	

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "The Pc chose rock"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "The Pc chose paper"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "The Pc chose scissors"
	}

	messageInt := rand.Intn(3)

	var message string

	if playerValue == computerValue {
		roundResult = "It's a draw"
		message = drawMesaages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "Player Wins"
		message = winMesaages[messageInt]
	} else {
		ComputerScore++
		roundResult = "PC Wins"	
		message = loseMesaages[messageInt]
	}

	result := Round{
		Message: message,
		ComputerChoice: computerChoice,
		RoundResult: roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore: strconv.Itoa(ComputerScore),
		PlayerScore: strconv.Itoa(PlayerScore),
	}

	return result

}