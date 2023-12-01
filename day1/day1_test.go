package day1

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {
	sum, err := Part1("./data/day1_test.txt")
	if err != nil {
		log.Fatalf("Error processing part 1: %s", err)
	}
	expected := 142
	assert.Equal(t, sum, expected, "Sum is %d but should be %d", sum, expected)
}

func TestDay1Part2(t *testing.T) {
	sum, err := Part2("./data/day1_part2_test.txt")
	if err != nil {
		log.Fatalf("Error processing part 2: %s", err)
	}
	expected := 281
	assert.Equal(t, sum, expected, "Sum is %d but should be %d", sum, expected)
}
