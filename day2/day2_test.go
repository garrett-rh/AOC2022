package day2

import (
	"aoc2022/helpers"
	"os"
	"strings"
	"testing"
)

func TestDay2(t *testing.T) {
	tmp, err := os.ReadFile("input.test")
	helpers.Check(err)
	testInput := strings.Split(string(tmp), "\n")

	tmp, err = os.ReadFile("input")
	helpers.Check(err)
	realInput := strings.Split(string(tmp), "\n")

	t.Run("Day 2a Test Input", func(t *testing.T) {

		got := Scorer(1, testInput)
		if got != 15 {
			t.Errorf("got %v, want 15", got)
		}
		t.Logf("%v", got)
	})

	t.Run("Day 2a Input", func(t *testing.T) {
		got := Scorer(1, realInput)
		t.Logf("%v", got)
	})

	t.Run("Day 2b Test Input", func(t *testing.T) {
		got := Scorer(2, testInput)
		if got != 12 {
			t.Errorf("got %v, want 15", got)
		}
		t.Logf("%v", got)
	})
	t.Run("Day 2b Input", func(t *testing.T) {

		got := Scorer(2, realInput)
		t.Logf("%v", got)
	})
}
