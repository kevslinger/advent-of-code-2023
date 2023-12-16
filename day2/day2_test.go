package day2

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) {
	file, err := os.Open("./data/day2_test.txt")
	if err != nil {
		log.Fatalf("Error reading Day 2 Part 1 File: %s", err)
	}
	defer file.Close()

	sum, err := Part1(file, GetMatchers())
	if err != nil {
		log.Fatalf("Error processing part 1: %s", err)
	}
	expected := 8
	assert.Equal(t, expected, sum, "Sum is %d but should be %d", sum, expected)
}

func TestDay2Part2(t *testing.T) {
	file, err := os.Open("./data/day2_test.txt")
	if err != nil {
		log.Fatalf("Error reading Day 2 Part 2 File: %s", err)
	}
	defer file.Close()

	sum, err := Part2(file, GetMatchers())
	if err != nil {
		log.Fatalf("Error processing Day 2: %s", err)
	}

	expected := 2286
	assert.Equal(t, expected, sum, "Sum is %d but should be %d", sum, expected)
}
