package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	towndata "swiss-town-guess-game.com/swiss-town-guess-game/swisspostdata"
)

const QUESTIONS_PER_QUIZ = 3

func main() {
	fmt.Println("**************************************************************************************************************")
	fmt.Println("* Welcome to the Swiss Town Guess Game!                                                                      *")
	fmt.Println("*                                                                                                            *")
	fmt.Println("* Thanks to post.ch / swisspost.opendatasoft.com / BFS for providing free data for this quiz.                *")
	fmt.Println("**************************************************************************************************************")

	fmt.Printf("Please enter your valid swisspost.opendatasoft.com API-Key: ")
	apiKeyInputValue, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println("Could not read api key")
		os.Exit(1)
	}
	fmt.Println()

	apiKey := string(apiKeyInputValue)

	for i := 1; i <= QUESTIONS_PER_QUIZ; i++ {

		// Get random number between 1 and 770
		// BFS numbers for Swiss towns are in that range
		bfsnr := towndata.GetRandomSwissTownBFSNumber()

		// Get data of town with bfsnr
		townInfo, err := towndata.GetTown(bfsnr, apiKey)
		if err != nil {
			fmt.Println("/// error at towndata.GetTown happened")
			fmt.Println(err)
			os.Exit(1)
		}

		// Print town info
		fmt.Println("//////// towninfo")
		fmt.Println(*townInfo)
	}
}
