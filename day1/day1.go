package day1

import (
	"strconv"
	"strings"
)

func FileConversion(file []byte) [][]int {
	tmp := strings.Split(string(file), "\n")

	var buf []int
	var final [][]int
	for _, line := range tmp {
		num, err := strconv.Atoi(line)

		if err != nil {
			final = append(final, buf)
			buf = []int{}
		} else {
			buf = append(buf, num)
		}
	}
	return final
}

func TopElf(input [][]int) (max int, idx int) {
	currElf := 0
	max = 0
	idx = 0

	for index, cals := range input {
		for _, item := range cals {
			currElf += item
		}
		if currElf > max {
			max = currElf
			idx = index
		}
		currElf = 0
	}
	return max, idx
}

func Top3Elf(input [][]int) int {
	elfOne, elfIDX := TopElf(input)
	remove(input, elfIDX)
	elfTwo, elfIDX := TopElf(input)
	remove(input, elfIDX)
	elfThree, _ := TopElf(input)

	return elfOne + elfTwo + elfThree
}

func remove(s [][]int, i int) [][]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
