package day2

func Outcome(part int, line string) int {
	//Part 1
	//column 1
	//A for rock
	//B for paper
	//C for scissors

	//column 2
	//X for rock
	//Y for paper
	//Z for scissors

	//scoring
	//+1 for rock
	//+2 for paper
	//+3 for scissors
	//+0 for loss
	//+3 for draw
	//+6 for win
	if part == 1 {
		outcomes := map[string]int{"A Y": 8, "A X": 4, "A Z": 3, "B Y": 5, "B X": 1, "B Z": 9, "C Y": 2, "C X": 7, "C Z": 6}
		return outcomes[line]
	}
	//Part 2
	//column 1
	//A for rock
	//B for paper
	//C for scissors

	//column 2
	//X for lose
	//Y for draw
	//Z for win

	//scoring
	//+1 for rock
	//+2 for paper
	//+3 for scissors
	//+0 for loss
	//+3 for draw
	//+6 for win
	if part == 2 {
		outcomes := map[string]int{"A Y": 4, "A X": 3, "A Z": 8, "B Y": 5, "B X": 1, "B Z": 9, "C Y": 6, "C X": 2, "C Z": 7}
		return outcomes[line]
	}

	return 0
}

func Scorer(part int, lines []string) int {
	var score int
	for _, line := range lines {
		score += Outcome(part, line)
	}

	return score
}
