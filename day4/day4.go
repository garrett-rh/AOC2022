package day4

import (
	"aoc2022/helpers"
	"strconv"
	"strings"
)

func PartOne(input []string) int {
	var ans int
	for _, line := range input {
		if line == "" {
			continue
		}
		assignments := parser(line)
		if contained(assignments) {
			ans++
		}
	}
	return ans
}

func PartTwo(input []string) int {
	var ans int
	for _, line := range input {
		if line == "" {
			continue
		}
		assignments := parser(line)
		if overlap(assignments) || contained(assignments) {
			ans++
		}
	}
	return ans
}

func parser(line string) []int {
	var ranges []int
	for _, assignments := range strings.Split(line, ",") {
		elf := strings.Split(assignments, "-")
		start, err := strconv.Atoi(elf[0])
		helpers.Check(err)
		finish, err := strconv.Atoi(elf[1])
		helpers.Check(err)
		ranges = append(ranges, start, finish)
	}
	// gives back an array of ints formatted where:
	// idx 0 is start for elf 1
	// idx 1 is end for elf 1
	// idx 2 is start for elf 2
	// idx 3 is end for elf 2
	// a bit gross but oh well
	return ranges
}

func contained(assignments []int) bool {
	if assignments[0] <= assignments[2] && assignments[1] >= assignments[3] {
		return true
	}

	if assignments[2] <= assignments[0] && assignments[3] >= assignments[1] {
		return true
	}

	return false
}

func overlap(assignments []int) bool {
	if assignments[0] >= assignments[2] && assignments[0] <= assignments[3] {
		return true
	}
	if assignments[1] >= assignments[2] && assignments[1] <= assignments[3] {
		return true
	}
	return false
}
