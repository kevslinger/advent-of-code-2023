package day13

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay13(path string) {
	number, err := runner.RunPart(path, Part1)
	if err != nil {
		fmt.Printf("Error with Day 13 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 13 Part 1: %d\n", number)
	}

	number, err = runner.RunPart(path, Part2)
	if err != nil {
		fmt.Printf("Error with Day 13 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 13 Part 2: %d\n", number)
	}
}

type Note struct {
	Rows []string
	Cols []string
}

func Part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	var notes []Note = ParseNotes(scanner)
	var sum int
	for _, note := range notes {
		sum += SummarizeNote(note, 0)
	}
	return sum, nil
}

func Part2(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	var notes []Note = ParseNotes(scanner)
	var sum int
	for _, note := range notes {
		sum += SummarizeNote(note, 1)
	}
	return sum, nil
}

func ParseNotes(scanner *bufio.Scanner) []Note {
	var notes []Note = make([]Note, 0)
	for scanner.Scan() {
		// Skip over empty lines
		for line := scanner.Text(); len(line) == 0; {
			scanner.Scan()
		}
		notes = append(notes, ParseRowsAndColumns(scanner))
	}

	return notes
}

func ParseRowsAndColumns(scanner *bufio.Scanner) Note {
	var rows []string = make([]string, 0)

	// The scanner is currently pointing at a line we want
	var line string = scanner.Text()
	rows = append(rows, line)
	var cols [][]string = make([][]string, len(line))
	for i := 0; i < len(scanner.Text()); i++ {
		cols[i] = make([]string, 0)
		cols[i] = append(cols[i], string(line[i]))
	}
	for scanner.Scan() {
		line := scanner.Text()
		// End of note
		if len(line) == 0 {
			break
		}
		rows = append(rows, line)
		for i := 0; i < len(cols); i++ {
			cols[i] = append(cols[i], string(line[i]))
		}
	}
	var joinedCols []string = make([]string, len(cols))
	for i, col := range cols {
		joinedCols[i] = strings.Join(col, "")
	}
	return Note{Rows: rows, Cols: joinedCols}
}

func SummarizeNote(note Note, tolerance int) int {
	var sum int = CheckReflectionWithTolerance(note.Cols, tolerance)
	if sum == -1 {
		sum = CheckReflectionWithTolerance(note.Rows, tolerance) * 100
	}
	return sum
}

func CheckReflectionWithTolerance(lines []string, tolerance int) int {
	// Don't check the last row
	for i := 0; i < len(lines)-1; i++ {
		var smudgeAllowance int = tolerance - NumDifferences(lines, i, i+1)
		// Go from here until the beginning/end
		if smudgeAllowance >= 0 {
			left := i - 1
			right := i + 2
			for left >= 0 && right < len(lines) && smudgeAllowance >= 0 {
				smudgeAllowance -= NumDifferences(lines, left, right)
				left--
				right++
			}
			// Check if we made it to the end
			if smudgeAllowance == 0 && (left < 0 || right >= len(lines)) {
				return i + 1
			}
		}
	}
	return -1
}

func NumDifferences(lines []string, idx1, idx2 int) int {
	var numDiffs int
	for i := 0; i < len(lines[idx2]); i++ {
		if lines[idx1][i] != lines[idx2][i] {
			numDiffs++
		}
	}
	return numDiffs
}
