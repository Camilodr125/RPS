package rps

import (
	"fmt"
	"testing"
)

func TestPlayRoundt(t *testing.T) {

	for i := 0; i < 3; i++ {
		round := PlayRound(i)

		switch i {
		case 0:
			fmt.Println("The Player Chose Rock")
		case 1:
			fmt.Println("The Player Chose Paper")
		case 2:
			fmt.Println("The Player Chose Scissors")
		}

		fmt.Printf("Message: %s\n", round.Message)
		fmt.Printf("ComputerChoice: %s\n", round.ComputerChoice)
		fmt.Printf("Round Result: %s\n", round.RoundResult)
		fmt.Printf("ComputerChoiceInt: %d\n", round.ComputerChoiceInt)
		fmt.Printf("ComputerScore: %s\n", round.ComputerScore)
		fmt.Printf("PlayerScore: %s\n", round.PlayerScore)
		fmt.Println("\n -----------------------------------------------")
	}	
}
	