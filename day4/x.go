package main

import (
	"aoc2023/utils"
	"fmt"
	"math"
	"strings"
)

var TestInput = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

type card struct {
	id     int
	wn, mn []int
}

func parse(input []string) []card {
	var cards []card
	for _, line := range input {
		id, numbers, _ := strings.Cut(line, ": ")
		wns, mns, _ := strings.Cut(numbers, "|")

		wn := utils.IntsOfString(wns)
		mn := utils.IntsOfString(mns)

		id = strings.TrimSpace(strings.TrimPrefix(id, "Card"))
		cards = append(cards, card{utils.Atoi(id), wn, mn})
	}
	return cards
}

func score(c card) int {
	wm := make(map[int]struct{})
	for _, w := range c.wn {
		wm[w] = struct{}{}
	}
	mul := 0
	for _, m := range c.mn {
		if _, ok := wm[m]; ok {
			mul++
		}
	}
	if mul < 2 {
		return mul
	}
	return int(math.Pow(float64(2), float64(mul-1)))
}

func silver(input []string) int {
	cards := parse(input)
	sum := 0
	for _, card := range cards {
		sum += score(card)
	}
	return sum
}

func count(c card) int {
	wm := make(map[int]struct{})
	for _, w := range c.wn {
		wm[w] = struct{}{}
	}
	cnt := 0
	for _, m := range c.mn {
		if _, ok := wm[m]; ok {
			cnt++
		}
	}
	return cnt
}

func gold(input []string) int {
	cards := parse(input)

	cbc := make(map[int][]struct{})
	for _, card := range cards {
		cbc[card.id] = append(cbc[card.id], struct{}{})
	}

	for _, card := range cards {
		cnt := count(card)
		for range cbc[card.id] {
			for i := card.id + 1; i <= cnt+card.id; i++ {
				cbc[i] = append(cbc[i], struct{}{})
			}
		}
	}

	sum := 0
	for _, v := range cbc {
		sum += len(v)
	}
	return sum
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput))
	fmt.Printf("SILVER: %d\n", silver(utils.Input("day4/input")))
	fmt.Printf("TEST GOLD: %d\n", gold(TestInput))
	fmt.Printf("GOLD: %d\n", gold(utils.Input("day4/input")))

}
