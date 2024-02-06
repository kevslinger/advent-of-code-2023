package day9

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9Part1(t *testing.T) {
	file, err := os.Open("./data/day9_test.txt")
	if err != nil {
		t.Fatalf("Error opening file in Day 9 Part 1: %s\n", err)
	}
	defer file.Close()

	sum, err := part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 9 Part 1: %s\n", err)
	}

	expected := 114
	assert.Equal(t, expected, sum, "Expected %d but got %d\n", expected, sum)
}

func TestDay9Part2(t *testing.T) {
	file, err := os.Open("./data/day9_test.txt")
	if err != nil {
		t.Fatalf("Error opening file in Day 9 Part 2: %s\n", err)
	}
	defer file.Close()

	sum, err := part2(file)
	if err != nil {
		t.Fatalf("Error processing Day 9 Part 2: %s\n", err)
	}

	expected := 2
	assert.Equal(t, expected, sum, "Expected %d but got %d\n", expected, sum)
}
