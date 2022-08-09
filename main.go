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
	apiKey, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println("Could not read api key")
		os.Exit(1)
	}
	fmt.Println()

	fmt.Println("Key entered: ")
	fmt.Println(string(apiKey))

	fmt.Println(towndata.HelloTest)
}
