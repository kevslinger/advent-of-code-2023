package day7

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7Part1(t *testing.T) {
	file, err := os.Open("./data/day7_test.txt")
	if err != nil {
		t.Fatalf("Error reading file in Day 7 Part 1: %s\n", err)
	}
	defer file.Close()

	winnings, err := Part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 7 Part 1: %s\n", err)
	}
	expected := 6440
	assert.Equal(t, expected, winnings, "Expected %d but got %d", expected, winnings)
}

func TestDay7Part2(t *testing.T) {
	file, err := os.Open("./data/day7_test.txt")
	if err != nil {
		t.Fatalf("Error reading file in Day 7 Part 1: %s\n", err)
	}
	defer file.Close()

	winnings, err := Part2(file)
	if err != nil {
		t.Fatalf("Error processing Day 7 Part 1: %s\n", err)
	}
	expected := 5905
	assert.Equal(t, expected, winnings, "Expected %d but got %d", expected, winnings)
}
