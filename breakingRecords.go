// breaking records [0, 1]

// looping {

// 	if currentscore != previousscore {
// 		if current score > previous score => breaking0 = breaking0  +
// 		if current score < previous score => breaking1 = breaking1 + 1

// 	}
// }

// mariaScore := []int32{12, 24, 10, 24}

package main

func breakingRecords(scores []int32) []int32 {
	var breakingRecords = []int32{0, 0}
	var currentScore int32
	var highestScore int32 = scores[0]
	var lowestScore int32 = scores[0]

	for _, v := range scores {

		currentScore = v

		if currentScore > lowestScore {
			breakingRecords[0]++
			lowestScore = currentScore
		}
		if currentScore < highestScore {
			breakingRecords[1]++
			highestScore = currentScore
		}

	}

	return breakingRecords
}
