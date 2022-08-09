package cantondata

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type CantonQuestionInfoSet struct {
	CantonQuestionInfos []CantonQuestionInfo
}

type CantonQuestionInfo struct {
	CantonAbbreviation string
	CorrectAnswer      bool
	CantonName         string
}

// Returns a shuffeled canton question info set.
func GetCantonQuestionInfoSet(correctCantonAbbreviation string, incorrectAnswerCount int) (cantonQuestionInfoSet *CantonQuestionInfoSet, err error) {

	var cantonQuestionInfos CantonQuestionInfoSet

	// Add correct canton
	var correctCantonInfo CantonQuestionInfo

	correctCantonInfo.CorrectAnswer = true
	correctCantonInfo.CantonAbbreviation = correctCantonAbbreviation
	correctCantonNamePointer := getCantonName(correctCantonAbbreviation)
	if correctCantonNamePointer == nil {
		return nil, errors.New(fmt.Sprintf("Canton name for correct canton abbreviation: %s could not be found!", correctCantonAbbreviation))
	}
	correctCantonName := *correctCantonNamePointer
	correctCantonInfo.CantonName = correctCantonName

	cantonQuestionInfos.CantonQuestionInfos = append(cantonQuestionInfos.CantonQuestionInfos, correctCantonInfo)

	// Add more random cantons (incorrect answers).
	incorrectCantonQuestionInfosPointer, err := getIncorrectCantonQuestionInfo(correctCantonName, incorrectAnswerCount)
	if incorrectCantonQuestionInfosPointer == nil {
		return nil, err
	}
	incorrectCantonQuestionInfos := *incorrectCantonQuestionInfosPointer
	cantonQuestionInfos.CantonQuestionInfos = append(cantonQuestionInfos.CantonQuestionInfos, incorrectCantonQuestionInfos...)

	// Shuffle cantons so that the correct entry is not always the first.
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cantonQuestionInfos.CantonQuestionInfos), func(i, j int) {
		cantonQuestionInfos.CantonQuestionInfos[i],
			cantonQuestionInfos.CantonQuestionInfos[j] = cantonQuestionInfos.CantonQuestionInfos[j],
			cantonQuestionInfos.CantonQuestionInfos[i]
	})

	return &cantonQuestionInfos, nil
}

func getIncorrectCantonQuestionInfo(correctCantonAbbreviation string, incorrectCantonAnswerCount int) (*[]CantonQuestionInfo, error) {

	var incorrectCantonInfos []CantonQuestionInfo

	takenCantonAbbreviations := []string{correctCantonAbbreviation}

	for i := 1; i <= incorrectCantonAnswerCount; i++ {

		// Get a canton that is not already taken
		for {
			cantonAlreadyTaken := false

			// Get random canton
			cantonQuestionInfoPointer, err := getRandomCanton()
			if err != nil {
				return nil, err
			}
			cantonQuestionInfo := *cantonQuestionInfoPointer

			// Check if canton is already taken
			for _, alreadyTakenCantonAbbreviation := range takenCantonAbbreviations {
				if cantonQuestionInfo.CantonAbbreviation == alreadyTakenCantonAbbreviation {
					cantonAlreadyTaken = true
				}
			}

			if !cantonAlreadyTaken {
				// add to result list
				incorrectCantonInfos = append(incorrectCantonInfos, cantonQuestionInfo)

				// add to taken list
				takenCantonAbbreviations = append(takenCantonAbbreviations, cantonQuestionInfo.CantonAbbreviation)

				// canton successfully added
				break
			}
		}
	}

	return &incorrectCantonInfos, nil
}
