package main

import (
	"aoc2023/util"
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
	cubes []cubes
	tr    int
	tg    int
	tb    int
}

func (g *game) add(c cubes) {
	g.cubes = append(g.cubes, c)
	g.tr = max(g.tr, c.red)
	g.tg = max(g.tg, c.green)
	g.tb = max(g.tb, c.blue)
}

func (g *game) check(cr, cg, cb int) bool {
	return g.tr <= cr && g.tg <= cg && g.tb <= cb
}

type cubes struct {
	red   int
	green int
	blue  int
}

func parseGame(line string) *game {
	parts := util.SplitTrim(line, ":")
	g := &game{
		id: util.Atoi(strings.TrimPrefix(parts[0], "Game ")),
	}

	bags := strings.Split(parts[1], ";")
	for _, bag := range bags {
		v := util.SplitTrim(bag, ",")
		c := cubes{}
		for _, vv := range v {
			pv := util.SplitTrim(vv, " ")
			if pv[1] == "green" {
				c.green = util.Atoi(pv[0])
			} else if pv[1] == "blue" {
				c.blue = util.Atoi(pv[0])
			} else {
				c.red = util.Atoi(pv[0])
			}
		}
		g.add(c)
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
		sum += g.tr * g.tg * g.tb
	}
	return sum
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput, 12, 13, 14))
	fmt.Printf("TEST GOLD: %d\n", gold(TestInput))

	fmt.Printf("SILVER: %d\n", silver(util.Input("./day2/input"), 12, 13, 14))
	fmt.Printf("GOLD: %d\n", gold(util.Input("./day2/input")))
}
