package day3

import (
	"aoc2022/helpers"
	"os"
	"strings"
	"testing"
)

func TestDay3(t *testing.T) {
	tmp, err := os.ReadFile("input.test")
	testInput := strings.Split(string(tmp), "\n")
	helpers.Check(err)

	tmp, err = os.ReadFile("input")
	realInput := strings.Split(string(tmp), "\n")
	helpers.Check(err)

	t.Run("Day 3a Test Input", func(t *testing.T) {

		got := PartOne(testInput)
		want := 157
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
		t.Logf("%v", got)
	})

	t.Run("Day 3a Input", func(t *testing.T) {
		t.Logf("%v", PartOne(realInput))
	})

	t.Run("Day 3b Test Input", func(t *testing.T) {
		got := PartTwo(testInput)
		want := 70
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
		t.Logf("%v", got)
	})
	t.Run("Day 3b Input", func(t *testing.T) {
		t.Logf("%v", PartTwo(realInput))
	})
}
