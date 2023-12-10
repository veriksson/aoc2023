package main

import (
	"aoc2023/utils"
	"fmt"
	"sort"
	"strings"
)

var TestInput = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

type race struct {
	time int
	dist int
}

func parse(input []string) []race {
	times := utils.IntsOfString(strings.TrimPrefix(input[0], "Time:"))
	dists := utils.IntsOfString(strings.TrimPrefix(input[1], "Distance:"))

	var races []race
	for i := 0; i < len(times); i++ {
		races = append(races, race{
			time: times[i],
			dist: dists[i],
		})
	}
	return races
}

func parseGold(input []string) race {
	time := utils.Atoi(strings.ReplaceAll(strings.TrimPrefix(input[0], "Time:"), " ", ""))
	dist := utils.Atoi(strings.ReplaceAll(strings.TrimPrefix(input[1], "Distance:"), " ", ""))

	return race{
		time: time,
		dist: dist,
	}
}

func speed(dtb, t int) int {
	steps := []int{}
	for i := 0; i < t; i++ {
		steps = append(steps, i)
	}

	a := sort.Search(len(steps), func(i int) bool {
		return steps[i]*(t-steps[i]) > dtb
	})
	lowest := steps[a]
	sort.Slice(steps, func(i, j int) bool {
		return steps[i] > steps[j]
	})

	b := sort.Search(len(steps), func(i int) bool {
		return steps[i]*(t-steps[i]) > dtb
	})
	highest := steps[b]

	return highest - lowest + 1
}

func silver(input []string) int {
	races := parse(input)
	sum := 1
	for _, race := range races {
		speed := speed(race.dist, race.time)
		sum *= speed
	}
	return sum
}

func gold(input []string) int {
	megarace := parseGold(input)
	return speed(megarace.dist, megarace.time)
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput))
	fmt.Printf("SILVER: %d\n", silver(utils.Input("day6/input")))
	fmt.Printf("TEST GOLD: %d\n", gold(TestInput))
	fmt.Printf("GOLD: %d\n", gold(utils.Input("day6/input")))

}
