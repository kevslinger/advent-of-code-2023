package day15

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay15Part1(t *testing.T) {
	file, err := os.Open("./data/day15_test.txt")
	if err != nil {
		t.Fatalf("Error opening input for Day 15 Part 1: %s\n", err)
	}
	defer file.Close()

	sum, err := part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 15 Part 1: %s\n", err)
	}
	expected := 1320
	assert.Equal(t, expected, sum, "Expected %d but got %d", expected, sum)
}

func TestDay15Part2(t *testing.T) {
	file, err := os.Open("./data/day15_test.txt")
	if err != nil {
		t.Fatalf("Error opening input for Day 15 Part 1: %s\n", err)
	}
	defer file.Close()

	power, err := part2(file)
	if err != nil {
		t.Fatalf("Error processing Day 15 Part 1: %s\n", err)
	}
	expected := 145
	assert.Equal(t, expected, power, "Expected %d but got %d", expected, power)
}
