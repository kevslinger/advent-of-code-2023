package day2

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) { testDay2Part1(t, "./data/day2_test.txt") }

func testDay2Part1(t *testing.T, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error reading Day 2 Part 1 File: %s", err)
	}
	defer file.Close()

	sum, err := part1(file)
	if err != nil {
		log.Fatalf("Error processing part 1: %s", err)
	}
	expected := 8
	assert.Equal(t, expected, sum, "Sum is %d but should be %d", sum, expected)
}

func TestDay2Part2(t *testing.T) { testDay2Part2(t, "./data/day2_test.txt") }

func testDay2Part2(t *testing.T, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error reading Day 2 Part 2 File: %s", err)
	}
	defer file.Close()

	sum, err := part2(file)
	if err != nil {
		log.Fatalf("Error processing Day 2: %s", err)
	}

	expected := 2286
	assert.Equal(t, expected, sum, "Sum is %d but should be %d", sum, expected)
}
