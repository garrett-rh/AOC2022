package day5

import (
	"aoc2022/helpers"
	"os"
	"testing"
)

func TestDay5(t *testing.T) {
	tmp, err := os.ReadFile("input.test")
	testInput := string(tmp)
	helpers.Check(err)

	tmp, err = os.ReadFile("input")
	realInput := string(tmp)
	helpers.Check(err)

	t.Run("Day 5a Test Input", func(t *testing.T) {

		got := PartOne(testInput)
		want := "CMZ"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
		t.Logf("%v", got)
	})

	t.Run("Day 5a Input", func(t *testing.T) {
		t.Logf("%v", PartOne(realInput))
	})

	t.Run("Day 5b Test Input", func(t *testing.T) {
		got := PartTwo(testInput)
		want := "MCD"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
		t.Logf("%v", got)
	})
	t.Run("Day 5b Input", func(t *testing.T) {
		t.Logf("%v", PartTwo(realInput))
	})
}
