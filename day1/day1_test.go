package day1

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) { testDay1Part1(t, "./data/day1_test.txt") }

func testDay1Part1(t *testing.T, path string) {
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Cannot open file in Day 1 Part 1: %s", err)
	}
	defer file.Close()

	sum, err := part1(file)
	if err != nil {
		t.Fatalf("Error processing part 1: %s", err)
	}
	expected := 142
	assert.Equal(t, expected, sum, "Sum is %d but should be %d", sum, expected)
}

func TestDay1Part2(t *testing.T) { testDay1Part2(t, "./data/day1_part2_test.txt") }

func testDay1Part2(t *testing.T, path string) {
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Cannot open file in Day 1 Part 2: %s", err)
	}
	defer file.Close()

	sum, err := part2(file)
	if err != nil {
		log.Fatalf("Error processing part 2: %s", err)
	}
	expected := 281
	assert.Equal(t, expected, sum, "Sum is %d but should be %d", sum, expected)
}
