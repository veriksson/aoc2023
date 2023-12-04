package main

import "fmt"
import "strings"
import "aoc2023/utils"

var TestInput = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

type card struct {
	wn, mn []int
}

func parse(input []string) []card {
	var cards []card
	for _, line := range input {
		_, numbers, _ := strings.Cut(line, ": ")
		wns, mns := strings.Cut(numbers, "|")

		wn := utils.IntsOfString(wns)
		mn := utils.IntsOfString(mns)

		cards = append(cards, card{wn, mn})
	}
	return cards
}

func score(c card) int {
	wm := make(map[int]struct{})
	s := 0
	for _, w := range c.wn {
		wm[w] = struct{}{}
	}
	for _, m := range c.mn {
		if _, ok := wm[m]; ok {
			s++
		}
	}
	return s
}

func silver(input []string) int {
	cards := parse(input)
	sum := 0
	for _, card := range cards {
		sum += score(card)
	}
	return sum
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput))
}
