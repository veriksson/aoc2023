package main

import (
	"aoc2023/utils"
	"fmt"
	"sort"
	"strings"
)

var TestInput = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

var TestOrderInput = []string{
	"AAA42 1",
	"AA96A 2",
}

var TestWeirdHouse = []string{
	"AK7AA 1",
	"AAJKK 2",
	"AA5A8 3",
}

var TestWeirdHouse2 = []string{
	"85J4J 1",
}

type hand struct {
	cards  []string
	values []int
	bid    int
}

const (
	TypeHighCard int = iota
	TypeOnePair
	TypeTwoPair
	TypeThreeOfAKind
	TypeFullhouse
	TypeFourOfAKind
	TypeFiveOfAKind
)

type kvp struct {
	key   string
	value int
}

func invert(cm map[string]int) []kvp {
	var ret []kvp
	for k, v := range cm {
		ret = append(ret, kvp{key: k, value: v})
	}
	return ret
}

var (
	valmap = map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	jokermap = map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 1,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
)

// UGLY
func score(h hand, jokers bool) int {

	cm := make(map[string]int)
	for _, c := range h.cards {
		cm[c]++
	}

	// if we have 5 jokers then early return
	if cm["J"] == 5 && jokers {
		return TypeFiveOfAKind
	}

	bonus := func() (get func(string) int) {
		jokerBonus := 0
		if jokers {
			jokerBonus = cm["J"]
			delete(cm, "J") // remove jokers since they will be used for others hands
		}

		get = func(s string) int {
			if s == "J" {
				return 0
			}
			current := jokerBonus
			jokerBonus = 0
			return current
		}

		return get
	}

	get := bonus()

	pairs := 0
	hasPair := false
	hasTwoPair := false
	hasThreeOfAKind := false
	hasFourOfAKind := false
	hasFiveOfAKind := false

	kvps := invert(cm)
	// sort the hand so we construct the best hands possible with jokers first
	sort.Slice(kvps, func(i, j int) bool {
		return kvps[i].value > kvps[j].value
	})

	for _, kvp := range kvps {
		v := kvp.value
		v += get(kvp.key)
		if v == 5 {
			hasFiveOfAKind = true
		} else if v == 4 {
			hasFourOfAKind = true
		} else if v == 3 {
			hasThreeOfAKind = true
		} else if v == 2 {
			pairs++
		}
	}

	if pairs == 1 {
		hasPair = true
	} else if pairs == 2 {
		hasTwoPair = true
	}

	if hasFiveOfAKind {
		return TypeFiveOfAKind
	} else if hasFourOfAKind {
		return TypeFourOfAKind
	} else if hasPair && hasThreeOfAKind {
		return TypeFullhouse
	} else if hasThreeOfAKind {
		return TypeThreeOfAKind
	} else if hasTwoPair {
		return TypeTwoPair
	} else if hasPair {
		return TypeOnePair
	} else {
		return TypeHighCard
	}
}

func compare(h1, h2 hand) bool {
	for i := range h1.values {
		h1c := h1.values[i]
		h2c := h2.values[i]
		if h1c > h2c {
			return false
		} else if h1c < h2c {
			return true
		}
	}
	return false
}

func parse(input []string, jokers bool) []hand {
	var hands []hand
	for _, line := range input {
		fs := strings.Fields(line)
		cards := strings.Split(fs[0], "")
		var values []int
		for _, c := range cards {
			if jokers {
				values = append(values, jokermap[c])
			} else {
				values = append(values, valmap[c])
			}
		}
		bid := utils.Atoi(fs[1])
		hands = append(hands, hand{
			cards:  cards,
			bid:    bid,
			values: values,
		})
	}
	return hands
}

func silver(input []string) int {
	cards := parse(input, false)
	sort.Slice(cards, func(i, j int) bool {
		s1 := score(cards[i], false)
		s2 := score(cards[j], false)
		if s1 == s2 {
			return compare(cards[i], cards[j])
		}
		return s1 < s2
	})

	sum := 0

	for i, c := range cards {
		// fmt.Printf("RANK %d: %v (%v)\n", i, c.cards, c.values)
		sum += (i + 1) * c.bid
	}

	return sum
}

func gold(input []string) int {
	cards := parse(input, true)
	sort.Slice(cards, func(i, j int) bool {
		s1 := score(cards[i], true)
		s2 := score(cards[j], true)
		if s1 == s2 {
			return compare(cards[i], cards[j])
		}
		return s1 < s2
	})

	sum := 0

	for i, c := range cards {
		// fmt.Printf("RANK  %d(%d): %v (%v) \n", i, score(c, true), c.cards, c.values)
		sum += (i + 1) * c.bid
	}

	return sum
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput))
	fmt.Printf("SILVER: %d\n", silver(utils.Input("day7/input")))
	fmt.Printf("TEST GOLD: %d\n", gold(TestInput))
	// fmt.Printf("TEST GOLD: %d\n", gold(TestWeirdHouse))
	// fmt.Printf("TEST GOLD: %d\n", gold(TestOrderInput))
	// fmt.Printf("TEST GOLD: %d\n", gold(TestWeirdHouse2))
	fmt.Printf("GOLD: %d\n", gold(utils.Input("day7/input")))
}
