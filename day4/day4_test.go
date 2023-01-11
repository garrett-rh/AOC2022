package day4

import (
	"aoc2022/helpers"
	"os"
	"strings"
	"testing"
)

func TestDay4(t *testing.T) {
	tmp, err := os.ReadFile("input.test")
	testInput := strings.Split(string(tmp), "\n")
	helpers.Check(err)

	tmp, err = os.ReadFile("input")
	realInput := strings.Split(string(tmp), "\n")
	helpers.Check(err)

	t.Run("Day 4a Test Input", func(t *testing.T) {

		got := PartOne(testInput)
		want := 4
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
		t.Logf("%v", got)
	})

	t.Run("Day 4a Input", func(t *testing.T) {
		t.Logf("%v", PartOne(realInput))
	})

	t.Run("Day 4b Test Input", func(t *testing.T) {
		got := PartTwo(testInput)
		want := 8
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
		t.Logf("%v", got)
	})
	t.Run("Day 4b Input", func(t *testing.T) {
		t.Logf("%v", PartTwo(realInput))
	})
}
