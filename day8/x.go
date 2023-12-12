package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

var TestInput1 = []string{
	"RL",
	"",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)",
}

var TestInput2 = []string{
	"LLR",
	"",
	"AAA = (BBB, BBB)",
	"BBB = (AAA, ZZZ)",
	"ZZZ = (ZZZ, ZZZ)",
}

var TestInput3 = []string{
	"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
}

type navigationmap struct {
	steps []string
	elem  map[string][]string
	start string
	count int
}

func (nm *navigationmap) next(from string) string {
	dir := nm.steps[nm.count%len(nm.steps)]

	nm.count++
	if dir == "L" {
		return nm.elem[from][0]
	} else {
		return nm.elem[from][1]

	}
}

func (nm *navigationmap) ghost(froms []string) ([]string, bool) {
	dir := nm.steps[nm.count%len(nm.steps)]
	nm.count++
	var nexts []string
	exited := true
	if dir == "L" {
		for _, from := range froms {
			next := nm.elem[from][0]
			if next[2] != 'Z' {
				exited = false
			}
			nexts = append(nexts, next)
		}
	} else {
		for _, from := range froms {
			next := nm.elem[from][1]
			if next[2] != 'Z' {
				exited = false
			}
			nexts = append(nexts, next)
		}
	}
	return nexts, exited
}

func ghostloop(start string, m map[string][]string, dirs []string) int {
	// fast := 0
	slow := 0
	curs := start
	// curf := start

	move := func(k, dir string) string {
		if dir == "L" {
			return m[k][0]
		} else {
			return m[k][1]
		}
	}
	for {
		nexts := dirs[slow%len(dirs)]
		curs = move(curs, nexts)
		slow++

		if curs[2] == 'Z' {
			return slow
		}

		// nextf := dirs[fast%len(dirs)]
		// curf = move(curf, nextf)
		// fast++

		// nextf = dirs[fast%len(dirs)]
		// curf = move(curf, nextf)
		// fast++

		// if curs[2] != 'Z' {
		// 	continue
		// }
		// if curs == curf {
		// 	return slow
		// }
	}
}

type elem struct {
	name  string
	left  *elem
	right *elem
}

func parse(input []string) *navigationmap {
	steps := strings.Split(input[0], "")

	m := make(map[string][]string)
	start := ""
	for i := 2; i < len(input); i++ {
		name, lr, _ := strings.Cut(input[i], " = ")
		if start == "" {
			start = name
		}
		l := lr[1:4]
		r := lr[6:9]
		m[name] = []string{l, r}
	}

	return &navigationmap{
		steps: steps,
		elem:  m,
		start: start,
		count: 0,
	}
}

func silver(input []string) int {
	nm := parse(input)
	current := "AAA"
	for {
		current = nm.next(current)
		if current == "ZZZ" {
			break
		}
	}
	return nm.count
}

func ghostStarts(nm *navigationmap) []string {
	var starts []string
	for k := range nm.elem {
		if k[2] == 'A' {
			starts = append(starts, k)
		}
	}
	return starts
}

func gold(input []string) int {
	nm := parse(input)
	currents := ghostStarts(nm)
	var rets []int
	for _, c := range currents {
		rets = append(rets, ghostloop(c, nm.elem, nm.steps))
	}
	a := rets[0]
	b := rets[1]
	return utils.LCM(a, b, rets[2:]...)
}

func main() {
	// fmt.Printf("TEST SILVER: %d\n", silver(TestInput1))
	// fmt.Printf("TEST SILVER: %d\n", silver(TestInput2))
	// fmt.Printf("TEST SILVER: %d\n", silver(utils.Input("day8/input")))
	fmt.Printf("TEST GOLD: %d\n", gold(TestInput3))
	fmt.Printf("GOLD: %d\n", gold(utils.Input("day8/input")))
}
