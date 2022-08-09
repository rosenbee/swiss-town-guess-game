package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	cantondata "swiss-town-guess-game.com/swiss-town-guess-game/cantondata"
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

		// User info that first question is being prepared
		fmt.Println(fmt.Sprintf("Preparing Question %d. Please wait.", i))

		// Get data of town with bfsnr
		var townInfo *towndata.TownInfo

		// Try to get a valid town until one is found
		// TODO: would be nice to have a max try counter
		for townInfo == nil {

			// We use the BFS number to query town data
			bfsnr := towndata.GetRandomSwissTownBFSNumber()

			// Get towninfo from post.ch api
			townInfo, err = towndata.GetTown(bfsnr, apiKey)
			if err != nil {
				fmt.Println("/// error at towndata.GetTown happened")
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// Print town info
		town := *townInfo
		cantonNamePointer := cantondata.GetCantonName(town.CantonCode)
		if cantonNamePointer == nil {
			fmt.Println(fmt.Sprintf("/// canton with abbreviation %s could not be found", town.CantonCode))
			os.Exit(1)
		}
		cantonName := *cantonNamePointer

		fmt.Println("//////// towninfo")
		fmt.Println(town)
		fmt.Println(cantonName)

	}
}
