package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

var TestInput = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

type game struct {
	id    int
	red   int
	green int
	blue  int
}

func (g *game) check(cr, cg, cb int) bool {
	return g.red <= cr && g.green <= cg && g.blue <= cb
}

func parseGame(line string) *game {
	parts := utils.SplitTrim(line, ":")
	g := &game{
		id: utils.Atoi(strings.TrimPrefix(parts[0], "Game ")),
	}

	rounds := strings.Split(parts[1], ";")
	for _, round := range rounds {
		vs := utils.SplitTrim(round, ",")
		for _, v := range vs {
			pv := strings.Fields(v)
			switch pv[1] {
			case "green":
				g.green = max(g.green, utils.Atoi(pv[0]))
			case "blue":
				g.blue = max(g.blue, utils.Atoi(pv[0]))
			case "red":
				g.red = max(g.red, utils.Atoi(pv[0]))
			}
		}
	}
	return g
}

func silver(input []string, cr, cg, cb int) int {
	var games []*game

	for _, l := range input {
		games = append(games, parseGame(l))
	}

	sum := 0
	for _, g := range games {
		if g.check(cr, cg, cb) {
			sum += g.id
		}
	}
	return sum
}

func gold(input []string) int {
	var games []*game

	for _, l := range input {
		games = append(games, parseGame(l))
	}

	sum := 0
	for _, g := range games {
		sum += g.red * g.green * g.blue
	}
	return sum
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput, 12, 13, 14))
	fmt.Printf("TEST GOLD: %d\n", gold(TestInput))

	fmt.Printf("SILVER: %d\n", silver(utils.Input("./day2/input"), 12, 13, 14))
	fmt.Printf("GOLD: %d\n", gold(utils.Input("./day2/input")))
}
