package day3

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) { testDay3Part1(t, "./data/day3_test.txt") }

func testDay3Part1(t *testing.T, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error processing file in Day 3: %s", err)
	}
	defer file.Close()

	sum, err := part1(file)
	if err != nil {
		log.Fatalf("Error with Day 3 Part 1: %s", err)
	}
	expected := 4361
	assert.Equal(t, expected, sum, "Sum is %d but should be %d", sum, expected)
}

func TestDay3Part2(t *testing.T) { testDay3Part2(t, "./data/day3_test.txt") }

func testDay3Part2(t *testing.T, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error processing file in Day 3: %s", err)
	}
	defer file.Close()

	sum, err := part2(file)
	if err != nil {
		log.Fatalf("Error with Day 3 Part 2: %s", err)
	}
	expected := 467835
	assert.Equal(t, expected, sum, "Sum is %d but should be %d", sum, expected)
}
