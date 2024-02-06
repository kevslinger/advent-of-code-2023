package day16

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay16Part1(t *testing.T) {
	file, err := os.Open("./data/day16_test.txt")
	if err != nil {
		t.Fatalf("Error opening file in Day 16 Part 1: %s\n", err)
	}
	defer file.Close()

	tiles, err := part1(file)
	if err != nil {
		t.Fatalf("Error wth Day 16 Part 1: %s\n", err)
	}
	var expected int = 46
	assert.Equal(t, expected, tiles, "Expected %d but got %d\n", expected, tiles)
}

func TestDay16Part2(t *testing.T) {
	file, err := os.Open("./data/day16_test.txt")
	if err != nil {
		t.Fatalf("Error opening file in Day 16 Part 1: %s\n", err)
	}
	defer file.Close()

	tiles, err := part2(file)
	if err != nil {
		t.Fatalf("Error wth Day 16 Part 1: %s\n", err)
	}
	var expected int = 51
	assert.Equal(t, expected, tiles, "Expected %d but got %d\n", expected, tiles)
}
