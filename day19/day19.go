package day19

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay19(path string) {
	rating, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error with Day 19 part 1: %s\n", err)
	} else {
		fmt.Printf("Answer for Day 19 part 1: %d\n", rating)
	}
}

type condition struct {
	category   byte
	operation  byte
	comparison int
	result     string
}

func (c condition) Check(p part) bool {
	if c.operation == '<' {
		return p.Values[c.category] < c.comparison
	} else {
		return p.Values[c.category] > c.comparison
	}
}

type workflow struct {
	name       string
	conditions []condition
	def        string
}

type part struct {
	Values map[byte]int
}

func (p part) GetRating() int {
	return p.Values['x'] + p.Values['m'] + p.Values['a'] + p.Values['s']
}

func part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	workflows, err := parseWorkflows(scanner)
	if err != nil {
		return -1, err
	}
	parts, err := parseParts(scanner)
	if err != nil {
		return -1, err
	}
	return computeRatings(workflows, parts), nil
}

func parseWorkflows(scanner *bufio.Scanner) (map[string]workflow, error) {
	workflows := make(map[string]workflow)
	for scanner.Scan() {
		line := scanner.Text()
		// Split between workflows and parts
		if len(line) == 0 {
			break
		}
		wflow, err := parseWorkflow(line)
		if err != nil {
			return make(map[string]workflow), err
		}
		workflows[wflow.name] = wflow
	}
	return workflows, nil
}

func parseWorkflow(line string) (workflow, error) {
	numberMatcher := regexp.MustCompile("[0-9]+")
	// Need to chop off }
	toks := strings.Split(line[:len(line)-1], "{")
	name := toks[0]
	conditionStrs := strings.Split(toks[1], ",")
	conditions := make([]condition, len(conditionStrs)-1)
	for idx, condStr := range conditionStrs[:len(conditionStrs)-1] {
		cat := condStr[0]
		op := condStr[1]
		compStr := numberMatcher.FindString(condStr)
		comparison, err := strconv.Atoi(compStr)
		if err != nil {
			return workflow{}, err
		}
		result := strings.Split(condStr, ":")[1]
		conditions[idx] = condition{category: cat, operation: op, comparison: comparison, result: result}
	}

	return workflow{name: name, conditions: conditions, def: conditionStrs[len(conditionStrs)-1]}, nil
}

func parseParts(scanner *bufio.Scanner) ([]part, error) {
	parts := make([]part, 0)
	for scanner.Scan() {
		line := scanner.Text()
		p, err := parsePart(line)
		if err != nil {
			return parts, err
		}
		parts = append(parts, p)
	}
	return parts, nil
}

func parsePart(line string) (part, error) {
	categories := strings.Split(line[1:len(line)-1], ",")
	catMap := make(map[byte]int)
	for _, cat := range categories {
		num, err := strconv.Atoi(cat[2:])
		if err != nil {
			return part{}, err
		}
		catMap[cat[0]] = num
	}
	return part{catMap}, nil
}

func computeRatings(workflows map[string]workflow, parts []part) int {
	var ratingSum int
	for _, p := range parts {
		ratingSum += ComputeRating(workflows, p)
	}
	return ratingSum
}

func ComputeRating(workflows map[string]workflow, p part) int {
	var result string = "in"
	for result != "A" && result != "R" {
		result = Processworkflow(workflows[result], p)
	}
	if result == "A" {
		return p.GetRating()
	} else {
		return 0
	}
}

func Processworkflow(workflow workflow, p part) string {
	for _, cond := range workflow.conditions {
		if cond.Check(p) {
			return cond.result
		}
	}
	return workflow.def
}
