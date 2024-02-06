package day8

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8Part1first(t *testing.T) {
	file, err := os.Open("./data/day8_test1.txt")
	if err != nil {
		t.Fatalf("Error opening file in Day 8 Part1_1: %s\n", err)
	}
	defer file.Close()

	steps, err := part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 8 Part 1_1: %s\n", err)
	}
	expected := 2
	assert.Equal(t, expected, steps, "Expected %d but got %d\n", expected, steps)
}

func TestDay8Part1second(t *testing.T) {
	file, err := os.Open("./data/day8_test2.txt")
	if err != nil {
		t.Fatalf("Error opening file in Day 8 Part1_1: %s\n", err)
	}
	defer file.Close()

	steps, err := part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 8 Part 1_1: %s\n", err)
	}
	expected := 6
	assert.Equal(t, expected, steps, "Expected %d but got %d\n", expected, steps)
}
