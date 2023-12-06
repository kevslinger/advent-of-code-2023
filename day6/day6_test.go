package day6

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6Part1(t *testing.T) {
	file, err := os.Open("./data/day6_test.txt")
	if err != nil {
		t.Fatalf("Error reading file: %s\n", err)
	}
	raceProduct, err := Part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 6 Part 1: %s\n", err)
	}
	expected := 288
	assert.Equal(t, expected, raceProduct, "Expected %d but got %d", expected, raceProduct)
}

func TestDay6Part2(t *testing.T) {
	file, err := os.Open("./data/day6_test.txt")
	if err != nil {
		t.Fatalf("Error reading file: %s\n", err)
	}
	raceProduct, err := Part2(file)
	if err != nil {
		t.Fatalf("Error processing Day 6 Part 1: %s\n", err)
	}
	expected := 71503
	assert.Equal(t, expected, raceProduct, "Expected %d but got %d", expected, raceProduct)
}
