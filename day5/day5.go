package day5

import (
	"aoc2022/helpers"
	"fmt"
	"strings"
)

type Stack []rune
type step struct {
	quantity int
	from     int
	to       int
}

func PartOne(input string) string {
	crates := crateParser(input)
	moves := moveParser(input)

	for _, move := range moves {
		for i := 0; i < move.quantity; i++ {
			top := crates[move.from][len(crates[move.from])-1]
			crates[move.to] = append(crates[move.to], top)
			crates[move.from] = crates[move.from][:len(crates[move.from])-1]
		}
	}
	var ans string
	for _, stack := range crates {
		ans += stack[len(stack)-1]
	}
	return ans
}

func PartTwo(input string) string {
	crates := crateParser(input)
	moves := moveParser(input)

	for _, move := range moves {
		fromIndex := len(crates[move.from]) - move.quantity
		crates[move.to] = append(crates[move.to], crates[move.from][fromIndex:]...)
		crates[move.from] = crates[move.from][:fromIndex]
	}

	var ans string
	for _, stack := range crates {
		ans += stack[len(stack)-1]
	}
	return ans
}

func moveParser(input string) []step {
	movesRaw := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	var moves []step
	for _, move := range movesRaw {
		tmp := step{}
		_, err := fmt.Sscanf(move, "move %d from %d to %d", &tmp.quantity, &tmp.from, &tmp.to)
		helpers.Check(err)

		tmp.from--
		tmp.to--
		moves = append(moves, tmp)
	}

	return moves
}

func crateParser(input string) [][]string {
	cargoRaw := strings.Split(input, "\n\n")[0]

	cargo := [][]string{}
	for _, row := range strings.Split(cargoRaw, "\n") {
		cargo = append(cargo, strings.Split(row, ""))
	}

	rows, cols := len(cargo), len(cargo[0])

	parsed := [][]string{}

	for c := 0; c < cols; c++ {
		if cargo[rows-1][c] != " " {
			stack := []string{}
			for r := rows - 2; r >= 0; r-- {
				char := cargo[r][c]
				if char != " " {
					stack = append(stack, char)
				}
			}
			parsed = append(parsed, stack)
		}
	}

	return parsed
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(char rune) {
	*s = append(*s, char)
}

func (s *Stack) pop() (rune, bool) {
	if s.isEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
