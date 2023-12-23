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
	rating, err := runner.RunPart(path, Part1)
	if err != nil {
		fmt.Printf("Error with Day 19 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer for Day 19 Part 1: %d\n", rating)
	}
}

type Condition struct {
	Category   byte
	Operation  byte
	Comparison int
	Result     string
}

func (c Condition) Check(part Part) bool {
	if c.Operation == '<' {
		return part.Values[c.Category] < c.Comparison
	} else {
		return part.Values[c.Category] > c.Comparison
	}
}

type Workflow struct {
	Name       string
	Conditions []Condition
	Default    string
}

type Part struct {
	Values map[byte]int
}

func (p Part) GetRating() int {
	return p.Values['x'] + p.Values['m'] + p.Values['a'] + p.Values['s']
}

func Part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	workflows, err := ParseWorkflows(scanner)
	if err != nil {
		return -1, err
	}
	parts, err := ParseParts(scanner)
	if err != nil {
		return -1, err
	}
	return ComputeRatings(workflows, parts), nil
}

func ParseWorkflows(scanner *bufio.Scanner) (map[string]Workflow, error) {
	workflows := make(map[string]Workflow)
	for scanner.Scan() {
		line := scanner.Text()
		// Split between workflows and parts
		if len(line) == 0 {
			break
		}
		workflow, err := ParseWorkflow(line)
		if err != nil {
			return make(map[string]Workflow), err
		}
		workflows[workflow.Name] = workflow
	}
	return workflows, nil
}

func ParseWorkflow(line string) (Workflow, error) {
	numberMatcher := regexp.MustCompile("[0-9]+")
	// Need to chop off }
	toks := strings.Split(line[:len(line)-1], "{")
	name := toks[0]
	conditionStrs := strings.Split(toks[1], ",")
	conditions := make([]Condition, len(conditionStrs)-1)
	for idx, condStr := range conditionStrs[:len(conditionStrs)-1] {
		cat := condStr[0]
		op := condStr[1]
		compStr := numberMatcher.FindString(condStr)
		comparison, err := strconv.Atoi(compStr)
		if err != nil {
			return Workflow{}, err
		}
		result := strings.Split(condStr, ":")[1]
		conditions[idx] = Condition{Category: cat, Operation: op, Comparison: comparison, Result: result}
	}

	return Workflow{Name: name, Conditions: conditions, Default: conditionStrs[len(conditionStrs)-1]}, nil
}

func ParseParts(scanner *bufio.Scanner) ([]Part, error) {
	parts := make([]Part, 0)
	for scanner.Scan() {
		line := scanner.Text()
		part, err := ParsePart(line)
		if err != nil {
			return parts, err
		}
		parts = append(parts, part)
	}
	return parts, nil
}

func ParsePart(line string) (Part, error) {
	categories := strings.Split(line[1:len(line)-1], ",")
	catMap := make(map[byte]int)
	for _, cat := range categories {
		num, err := strconv.Atoi(cat[2:])
		if err != nil {
			return Part{}, err
		}
		catMap[cat[0]] = num
	}
	return Part{catMap}, nil
}

func ComputeRatings(workflows map[string]Workflow, parts []Part) int {
	var ratingSum int
	for _, part := range parts {
		ratingSum += ComputeRating(workflows, part)
	}
	return ratingSum
}

func ComputeRating(workflows map[string]Workflow, part Part) int {
	var result string = "in"
	for result != "A" && result != "R" {
		result = ProcessWorkflow(workflows[result], part)
	}
	if result == "A" {
		return part.GetRating()
	} else {
		return 0
	}
}

func ProcessWorkflow(workflow Workflow, part Part) string {
	for _, cond := range workflow.Conditions {
		if cond.Check(part) {
			return cond.Result
		}
	}
	return workflow.Default
}
