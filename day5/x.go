package main

import (
	"aoc2023/utils"
	"fmt"
	"math"
	"strings"
)

var TestInput = []string{
	"seeds: 79 14 55 13",
	"",
	"seed-to-soil map:",
	"50 98 2",
	"52 50 48",
	"",
	"soil-to-fertilizer map:",
	"0 15 37",
	"37 52 2",
	"39 0 15",
	"",
	"fertilizer-to-water map:",
	"49 53 8",
	"0 11 42",
	"42 0 7",
	"57 7 4",
	"",
	"water-to-light map:",
	"88 18 7",
	"18 25 70",
	"",
	"light-to-temperature map:",
	"45 77 23",
	"81 45 19",
	"68 64 13",
	"",
	"temperature-to-humidity map:",
	"0 69 1",
	"1 0 69",
	"",
	"humidity-to-location map:",
	"60 56 37",
	"56 93 4",
}

type category struct {
	name   string
	ranges []catrange
	next   *category
}

type catrange struct {
	dst int
	src int
	len int
}

func (c *category) mapval(i int) int {
	for _, r := range c.ranges {
		if i < r.src || i > r.src+r.len {
			continue
		}
		return (i - r.src) + r.dst
	}
	return i
}

type almanac struct {
	seeds      []int
	categories *category
}

type step struct {
	val  int
	name string
	next *step
}

func parse(input []string) almanac {
	seeds := utils.IntsOfString(strings.TrimPrefix(input[0], "seeds: "))

	var categories []category
	for i := 1; i < len(input); i++ {
		if input[i] == "" {
			i++
		}

		name := strings.TrimSuffix(input[i], " map:")
		i++
		var ranges []catrange
		for {
			if i >= len(input) || input[i] == "" {
				categories = append(categories, category{
					name:   name,
					ranges: ranges,
				})
				break
			}
			ints := utils.IntsOfString(input[i])
			ranges = append(ranges, catrange{
				dst: ints[0],
				src: ints[1],
				len: ints[2],
			})
			i++
		}
	}

	first := &categories[0]
	current := first
	for i := 1; i < len(categories); i++ {
		current.next = &categories[i]
		current = &categories[i]
	}
	return almanac{
		seeds:      seeds,
		categories: first,
	}
}

func lookup(s *step, c *category) step {
	v := c.mapval(s.val)
	s2 := &step{
		name: c.name,
		val:  v,
	}
	s.next = s2
	if c.next != nil {
		lookup(s2, c.next)
	}
	return *s
}

func run(a almanac) []int {
	var steps []step
	for _, s := range a.seeds {
		t := &step{
			val:  s,
			name: "seed",
		}
		steps = append(steps, lookup(t, a.categories))
	}

	var lowest []int
	for _, s := range steps {
		for s.next != nil {
			s = *s.next
		}
		lowest = append(lowest, s.val)
	}

	return lowest
}

type seedrange struct {
	start int
	len   int
}

func constructRanges(seeds []int) []seedrange {
	var ret []seedrange
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		length := seeds[i+1]
		ret = append(ret, seedrange{
			start: start,
			len:   length,
		})
	}
	return ret
}

func run2(a almanac) int {
	sr := constructRanges(a.seeds)
	// sr = mergeRanges(sr)
	var best *step
	seen := make(map[int]struct{})
	for _, r := range sr {
		for i := r.start; i < r.start+r.len; i++ {
			if _, found := seen[i]; found {
				continue
			}
			s := &step{
				val:  i,
				name: "seed",
			}
			n := lookup(s, a.categories)
			for n.next != nil {
				n = *n.next
			}
			if best == nil || best.val > n.val {
				if best != nil {
					seen[best.val] = struct{}{}
				}
				best = &n
			}
		}
	}

	return best.val
}

func silver(input []string) int {
	almanac := parse(input)
	alts := run(almanac)

	sum := math.MaxInt
	for _, i := range alts {
		sum = min(sum, i)
	}
	return sum
}

func gold(input []string) int {
	almanac := parse(input)
	return run2(almanac)
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput))
	//	fmt.Printf("SILVER: %d\n", silver(utils.Input("day5/input")))
	fmt.Printf("TEST GOLD: %d\n", gold(TestInput))
	fmt.Printf("GOLD: %d\n", gold(utils.Input("day5/input")))
}
