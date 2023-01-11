package day1

import (
	"aoc2022/helpers"
	"os"
	"testing"
)

func TestDay1(t *testing.T) {
	testInput, err := os.ReadFile("input.test")
	helpers.Check(err)

	realInput, err := os.ReadFile("input")
	helpers.Check(err)

	t.Run("Day 1a Test Input", func(t *testing.T) {

		got, _ := TopElf(FileConversion(testInput))
		if got != 24000 {
			t.Errorf("got %v, want 24000", got)
		}
		t.Logf("%v", got)
	})

	t.Run("Day 1a Input", func(t *testing.T) {
		got, _ := TopElf(FileConversion(realInput))
		t.Logf("%v", got)
	})

	t.Run("Day 1b Test Input", func(t *testing.T) {

		got := Top3Elf(FileConversion(testInput))
		if got != 45000 {
			t.Errorf("got %v, want 45000", got)
		}
		t.Logf("%v", got)
	})
	t.Run("Day 1b Input", func(t *testing.T) {

		got := Top3Elf(FileConversion(realInput))
		t.Logf("%v", got)
	})
}
