package day3

import (
	"strings"
)

func runeToPriority(item rune) int {
	if item >= 65 && item <= 90 {
		item -= 38
	} else {
		item -= 96
	}

	return int(item)
}

func splitter(line string) (bagOne string, bagTwo string) {

	bagOne = line[0 : len(line)/2]
	bagTwo = line[len(line)/2:]

	return bagOne, bagTwo
}

func searcher(bagOne string, bagTwo string) int {
	var priority int
	for _, char := range bagOne {
		found := strings.Index(bagTwo, string(char))
		if found != -1 {
			priority = runeToPriority(char)
			break
		}
	}
	return priority
}

func PartOne(input []string) int {
	var prioritySum int
	for _, line := range input {
		prioritySum += searcher(splitter(line))
	}

	return prioritySum
}

func badgeLocator(group []string) int {
	var priority int
	for _, char := range group[0] {
		for _, item := range group[1] {
			foundOne := strings.Index(group[2], string(char))
			foundTwo := strings.Index(group[2], string(item))
			if foundOne == foundTwo && foundOne != -1 {
				priority = runeToPriority(char)
			}
		}
	}
	return priority
}

func PartTwo(input []string) int {
	var prioritySum int
	for idx := range input {
		var group []string
		if (idx+1)%3 == 0 {
			group = append(group, input[idx-2], input[idx-1], input[idx])
			prioritySum += badgeLocator(group)
		}
	}

	return prioritySum
}
