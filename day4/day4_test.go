package day4

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4Part1(t *testing.T) {
	file, err := os.Open("./data/day4_test.txt")
	if err != nil {
		t.Fatalf("Error reading file for Day 4 Part 1: %s", err)
	}
	defer file.Close()

	points, err := Part1(file)
	if err != nil {
		t.Fatalf("Error with processing Day 4 Part 1: %s", err)
	}
	expected := 13
	assert.Equal(t, expected, points, "Expected %d but got %d", expected, points)
}

func TestDay4Part2(t *testing.T) {
	file, err := os.Open("./data/day4_test.txt")
	if err != nil {
		t.Fatalf("Error reading file for Day 4 Part 1: %s", err)
	}
	defer file.Close()

	points, err := Part2(file)
	if err != nil {
		t.Fatalf("Error with processing Day 4 Part 1: %s", err)
	}
	expected := 30
	assert.Equal(t, expected, points, "Expected %d but got %d", expected, points)
}
