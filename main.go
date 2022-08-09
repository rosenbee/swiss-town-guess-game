package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	towndata "swiss-town-guess-game.com/swiss-town-guess-game/swisspostdata"
)

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

	fmt.Println("Key entered: ")
	fmt.Println(apiKey)

	fmt.Println(towndata.HelloTest)

	_, err = towndata.GetTown(1, apiKey)
	if err != nil {
		fmt.Println("/// error at towndata.GetTown happened")
		fmt.Println(err)
	}
}
