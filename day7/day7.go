package day7

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay7(path string) {
	winnings, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error procesing Day 7 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 7 Part 1: %d\n", winnings)
	}

	winnings, err = runner.RunPart(path, part2)
	if err != nil {
		fmt.Printf("Error processing Day 7 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 7 Part 2: %d\n", winnings)
	}
}

func part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	hands := make([]hand, 0)
	for scanner.Scan() {
		line := scanner.Text()
		hand, err := newHand(line)
		if err != nil {
			return -1, err
		}
		hands = append(hands, hand)
	}
	sort.Sort(byRank(hands))
	winnings := 0
	for idx, hand := range hands {
		winnings += (idx + 1) * hand.Bid
	}
	return winnings, nil
}

func part2(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	hands := make([]hand, 0)
	for scanner.Scan() {
		line := scanner.Text()
		hand, err := newHand(line)
		if err != nil {
			return -1, err
		}
		hand.updateRankWithJoker()
		hands = append(hands, hand)
	}
	sort.Sort(byRank(hands))
	winnings := 0
	for idx, hand := range hands {
		winnings += (idx + 1) * hand.Bid
	}
	return winnings, nil
}

type hand struct {
	Cards []int
	Rank  int
	Bid   int
}

func newHand(line string) (hand, error) {
	toks := strings.Split(line, " ")
	// toks[0] is hand, toks[1] is bid
	handSlice, err := parseHand(toks[0])
	if err != nil {
		return hand{}, err
	}
	rank := parseHandRank(handSlice)
	bid, err := parseBid(toks[1])
	if err != nil {
		return hand{}, err
	}
	return hand{handSlice, rank, bid}, nil
}

func parseHand(handStr string) ([]int, error) {
	hand := make([]int, 5)
	for idx, card := range handStr {
		switch card {
		case 'A':
			hand[idx] = 14
		case 'K':
			hand[idx] = 13
		case 'Q':
			hand[idx] = 12
		case 'J':
			hand[idx] = 11
		case 'T':
			hand[idx] = 10
		default:
			num, err := strconv.Atoi(string(card))
			if err != nil {
				return nil, err
			}
			hand[idx] = num
		}
	}
	return hand, nil
}

func parseHandRank(handSlice []int) int {
	m := make(map[int]int, 0)
	for _, card := range handSlice {
		m[card]++
	}
	var handRank int
	for _, v := range m {
		switch v {
		// 4/5 of a kind cannot have other options
		case 5:
			return 6
		case 4:
			return 5
		case 3:
			handRank += 3
		case 2:
			handRank += 1
		}
	}
	return handRank
}

func parseBid(bidStr string) (int, error) {
	return strconv.Atoi(bidStr)
}

func (h *hand) updateRankWithJoker() {
	m := make(map[int]int, 0)
	for idx, card := range h.Cards {
		if card == 11 {
			// For sorting
			h.Cards[idx] = 0
		}
		m[card]++
	}
	// Cannot improve if hand has no Jokers or is already 5oak
	if h.Rank >= 6 || m[11] <= 0 {
		return
	}
	switch h.Rank {
	// a full house with jokers being either side -> 5oak
	// 4 of a kind + jokers -> 5oak
	case 4, 5:
		if m[11] > 0 {
			h.Rank = 6
		}
	// 3 of a kind and any number of jokers -> 4oak
	case 3:
		if m[11] > 0 {
			h.Rank = 5
		}
	// 1 pair + a pair of jokers -> 4oak
	// 2 pair and a joker -> full house
	case 2:
		if m[11] == 2 {
			h.Rank = 5
		} else if m[11] == 1 {
			h.Rank = 4
		}
	// Pair of jokers or Pair + joker -> 3oak
	case 1:
		if m[11] > 0 {
			h.Rank = 3
		}
	// Joker -> pair
	case 0:
		if m[11] > 0 {
			h.Rank = 1
		}
	}
}

type byRank []hand

func (r byRank) Len() int      { return len(r) }
func (r byRank) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r byRank) Less(i, j int) bool {
	if r[i].Rank != r[j].Rank {
		return r[i].Rank < r[j].Rank
	}
	for k := 0; k < len(r[i].Cards); k++ {
		if r[i].Cards[k] == r[j].Cards[k] {
			continue
		}
		return r[i].Cards[k] < r[j].Cards[k]
	}
	return true
}
