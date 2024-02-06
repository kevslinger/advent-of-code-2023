package day13

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay13Part1(t *testing.T) {
	file, err := os.Open("./data/day13_test.txt")
	if err != nil {
		t.Fatalf("Cannot open file for Day 13 Part 1: %s\n", err)
	}
	defer file.Close()
	actual, err := part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 13 Part 1: %s\n", err)
	}
	expected := 405
	assert.Equal(t, expected, actual, "Expected %d but got %d\n", expected, actual)
}

func TestDay13Part2(t *testing.T) {
	file, err := os.Open("./data/day13_test.txt")
	if err != nil {
		t.Fatalf("Cannot open file for Day 13 Part 2: %s\n", err)
	}
	defer file.Close()
	actual, err := part2(file)
	if err != nil {
		t.Fatalf("Error processing Day 13 Part 2: %s\n", err)
	}
	expected := 400
	assert.Equal(t, expected, actual, "Expected %d but got %d\n", expected, actual)
}
