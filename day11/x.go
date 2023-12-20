package main

import (
	"aoc2023/utils"
	"fmt"
	"math"
	"sort"
)

var TestInput = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

type coord struct {
	x, y int
}

func manhattan(p1, p2 coord) int {
	return int(math.Abs(float64(p1.x-p2.x))) +
		int(math.Abs(float64(p1.y-p2.y)))
}

type galaxymap struct {
	runes    [][]rune
	galaxies []*coord
}

func expand(gm galaxymap) galaxymap {
	ymap := make(map[int]struct{})
	xmap := make(map[int]struct{})

	for _, g := range gm.galaxies {
		ymap[g.y] = struct{}{}
		xmap[g.x] = struct{}{}
	}

	yex := make(map[int]struct{})
	for y := 0; y < len(gm.runes); y++ {
		if _, s := ymap[y]; !s {
			yex[y] = struct{}{}
		}

	}
	xex := make(map[int]struct{})
	for x := 0; x < len(gm.runes[0]); x++ {
		if _, s := xmap[x]; !s {
			xex[x] = struct{}{}
		}
	}

	ylc := len(gm.runes[0]) + len(xex)
	exline := make([]rune, len(gm.runes[0])+len(xex))
	for a := 0; a < ylc; a++ {
		exline[a] = 'E'
	}

	var runes2 [][]rune
	for y := 0; y < len(gm.runes); y++ {
		var line []rune

		if _, ok := yex[y]; ok {
			runes2 = append(runes2, exline)
		}
		for x := 0; x < len(gm.runes[y]); x++ {
			line = append(line, gm.runes[y][x])
			if _, ok := xex[x]; ok {
				line = append(line, 'e')
			}
		}
		runes2 = append(runes2, line)
	}

	var galaxies []*coord
	for y := 0; y < len(runes2); y++ {
		for x := 0; x < len(runes2[y]); x++ {
			r := rune(runes2[y][x])
			if r == '#' {
				galaxies = append(galaxies, &coord{x, y})
			}
		}
	}
	return galaxymap{runes: runes2, galaxies: galaxies}
}

func expand2(gm galaxymap, factor int) galaxymap {
	ymap := make(map[int]struct{})
	xmap := make(map[int]struct{})

	for _, g := range gm.galaxies {
		ymap[g.y] = struct{}{}
		xmap[g.x] = struct{}{}
	}

	var yex []int
	for y := 0; y < len(gm.runes); y++ {
		if _, s := ymap[y]; !s {
			yex = append(yex, y)
		}
	}

	var xex []int
	for x := 0; x < len(gm.runes[0]); x++ {
		if _, s := xmap[x]; !s {
			xex = append(xex, x)
		}
	}

	update := func(f func(*coord)) {
		for i := 0; i < len(gm.galaxies); i++ {
			f(gm.galaxies[i])
		}
	}

	sort.Ints(yex)
	sort.Ints(xex)
	coordfactory := make(map[*coord]int)
	for _, k := range yex {
		update(func(c *coord) {
			if c.y > k {
				coordfactory[c] += factor
			}
		})
	}
	coordfactorx := make(map[*coord]int)
	for _, k := range xex {
		update(func(c *coord) {
			if c.x > k {
				coordfactorx[c] += factor
			}
		})
	}
	update(func(c *coord) {
		c.x += coordfactorx[c]
		c.y += coordfactory[c]
	})
	return gm
}

func printmap(gm galaxymap) {
	fmt.Println()

	for y := 0; y < len(gm.runes); y++ {
		for x := 0; x < len(gm.runes[y]); x++ {
			fmt.Printf("%c", gm.runes[y][x])
		}
		fmt.Println()
	}
	fmt.Println()

}

func printgalaxies(gm galaxymap) {
	fmt.Println()
	for i, g := range gm.galaxies {
		fmt.Printf("%d: (%d, %d)\n", i, g.x, g.y)
	}
	fmt.Println()
}

func pairs(gm galaxymap) [][2]*coord {
	var pairs [][2]*coord

	for i := 0; i < len(gm.galaxies)-1; i++ {
		for j := i + 1; j < len(gm.galaxies); j++ {
			pairs = append(pairs, [2]*coord{
				gm.galaxies[i],
				gm.galaxies[j],
			})
		}
	}
	return pairs
}

func parse(input []string) galaxymap {
	var runes [][]rune
	var galaxies []*coord
	for y := 0; y < len(input); y++ {
		var line []rune
		for x := 0; x < len(input[y]); x++ {
			r := rune(input[y][x])
			line = append(line, r)
			if r == '#' {
				galaxies = append(galaxies, &coord{x, y})
			}
		}
		runes = append(runes, line)
	}
	return galaxymap{runes, galaxies}
}

func silver(input []string) int {
	gm := parse(input)
	expanded := expand(gm)
	pairs := pairs(expanded)
	sum := 0
	for _, pair := range pairs {
		sum += manhattan(*pair[0], *pair[1])
	}
	return sum
}

func gold(input []string) int {
	gm := parse(input)
	expanded := expand2(gm, 999_999)
	pairs := pairs(expanded)
	sum := 0
	for _, pair := range pairs {
		sum += manhattan(*pair[0], *pair[1])
	}
	return sum
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput))
	fmt.Printf("SILVER: %d\n", silver(utils.Input("day11/input")))

	fmt.Printf("TEST GOLD: %d\n", gold(TestInput))
	fmt.Printf("GOLD: %d\n", gold(utils.Input("day11/input")))

}
