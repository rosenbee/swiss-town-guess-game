package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/patrickbucher/drawban"
	"golang.org/x/crypto/ssh/terminal"
	cantondata "swiss-town-guess-game.com/swiss-town-guess-game/cantondata"
	towndata "swiss-town-guess-game.com/swiss-town-guess-game/swisspostdata"
)

const (
	QUESTIONS_PER_QUIZ             = 5
	INCORRECT_ANSWERS_PER_QUESTION = 2
)

func main() {
	welcomeBanner := drawban.DrawBanner([]string{
		"Welcome to the Swiss Town Guess Game!",
		"Thanks to post.ch / swisspost.opendatasoft.com / BFS for providing free data for this quiz.",
	}, '*')
	fmt.Println(strings.Join(welcomeBanner, "\n"))

	var apiKey string
	flag.StringVar(&apiKey, "apikey", "", "API Key for swisspost.opendatasoft.com")
	flag.Parse()

	if apiKey == "" {
		fmt.Printf("Please enter your valid swisspost.opendatasoft.com API-Key: ")
		apiKeyInput, err := terminal.ReadPassword(0)
		if err != nil {
			fmt.Println("Could not read api key")
			os.Exit(1)
		}
		fmt.Println()
		apiKey = string(apiKeyInput)
	}

	reader := bufio.NewReader(os.Stdin)

	// Score
	var correctAnswers, incorrectAnswers int
	totalAnswers := INCORRECT_ANSWERS_PER_QUESTION + 1

	for i := 1; i <= QUESTIONS_PER_QUIZ; i++ {

		// User info that first question is being prepared
		fmt.Println("++++++++++++++++++++++++")
		fmt.Printf("Preparing Question %d. Please wait...\n", i)

		// Get data of town with bfsnr
		var townInfo *towndata.TownInfo

		// Try to get a valid town until one is found
		// TODO: would be nice to have a max try counter
		for townInfo == nil {

			// We use the BFS number to query town data
			bfsnr := towndata.GetRandomSwissTownBFSNumber()

			// Get towninfo from post.ch api
			townInfoInput, err := towndata.GetTown(bfsnr, apiKey)
			if err != nil {
				fmt.Println("/// error at towndata.GetTown happened")
				fmt.Println(err)
				os.Exit(1)
			}
			townInfo = townInfoInput
		}
		town := *townInfo

		cantonQuestionInfoSetPointer, err := cantondata.GetCantonQuestionInfoSet(town.CantonCode, INCORRECT_ANSWERS_PER_QUESTION)
		if err != nil {
			fmt.Println("Could not generate canton question set!")
			os.Exit(1)
		}
		cantonQuestionInfoSet := *cantonQuestionInfoSetPointer

		// Print question
		fmt.Printf("In which canton is %s?\n", town.Name)

		// Print possible answers
		for i, cantonInfo := range cantonQuestionInfoSet.CantonQuestionInfos {
			fmt.Printf("Type %d for %s\n", (i + 1), cantonInfo.CantonName)
		}

		// Await answer
		fmt.Print("Your answer: ")
		answerString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input.", err)
			os.Exit(1)
		}

		// Remove the delimiter from input string
		answerString = strings.TrimSuffix(answerString, "\n")

		// Check if answer is a valid number
		answer, err := strconv.Atoi(answerString)
		if err != nil || answer < 1 || answer > totalAnswers {
			fmt.Printf("Type in a number from 1 to %d\n", totalAnswers)
			fmt.Print("Your answer: ")
			answerString, err = reader.ReadString('\n')
			if err != nil || answer < 1 || answer > totalAnswers {
				fmt.Println("Sorry, you are not qualified for this game. Good bye!")
				os.Exit(1)
			}
		}

		// Check if answer is correct
		chosenAnswer := cantonQuestionInfoSet.CantonQuestionInfos[answer-1] // answer - 1, for index format
		if chosenAnswer.CorrectAnswer {
			fmt.Println("THAT WAS CORRECT!")
			correctAnswers++
		} else {
			fmt.Println("SORRY, INCORRECT")
			incorrectAnswers++
		}
		printScore(correctAnswers, incorrectAnswers)
	}

	goodbyeBanner := drawban.DrawBanner([]string{
		"This game is over. Thank you for playing!",
		fmt.Sprintf("You've got %d out of %d questions right!", correctAnswers, QUESTIONS_PER_QUIZ),
	}, '*')
	fmt.Println(strings.Join(goodbyeBanner, "\n"))
}

func printScore(correctAnswers, incorrectAnswers int) {
	fmt.Printf("Your new Score is: %d correct answers, %d incorrect answers.\n", correctAnswers, incorrectAnswers)
}
