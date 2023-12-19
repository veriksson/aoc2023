package main

import (
	"aoc2023/utils"
	"fmt"
)

var SimpleTestInput = []string{
	".....",
	".S-7.",
	".|.|.",
	".L-J.",
	".....",
}

var MessyTestInput = []string{
	"..F7.",
	".FJ|.",
	"SJ.L7",
	"|F--J",
	"LJ...",
}

var GoldTestInput1 = []string{
	"...........",
	".S-------7.",
	".|F-----7|.",
	".||.....||.",
	".||.....||.",
	".|L-7.F-J|.",
	".|..|.|..|.",
	".L--J.L--J.",
	"...........",
}

var GoldTestInput2 = []string{
	"..........",
	".S------7.",
	".|F----7|.",
	".||....||.",
	".||....||.",
	".|L-7F-J|.",
	".|..||..|.",
	".L--JL--J.",
	"..........",
}

var GoldTestInput3 = []string{
	"FF7FSF7F7F7F7F7F---7",
	"L|LJ||||||||||||F--J",
	"FL-7LJLJ||||||LJL-77",
	"F--JF--7||LJLJ7F7FJ-",
	"L---JF-JLJ.||-FJLJJ7",
	"|F|F-JF---7F7-L7L|7|",
	"|FFJF7L7F-JF7|JL---7",
	"7-L-JL7||F7|L7F-7F7|",
	"L.L7LFJ|||||FJL7||LJ",
	"L7JLJL-JLJLJL--JLJ.L",
}

var GoldTestInput4 = []string{
	".F----7F7F7F7F-7....",
	".|F--7||||||||FJ....",
	".||.FJ||||||||L7....",
	"FJL7L7LJLJ||LJ.L-7..",
	"L--J.L7...LJS7F-7L7.",
	"....F-J..F7FJ|L7L7L7",
	"....L7.F7||L7|.L7L7|",
	".....|FJLJ|FJ|F7|.LJ",
	"....FJL-7.||.||||...",
	"....L---J.LJ.LJLJ...",
}

var (
	WE = []int{0, -1, 0, 1} // horizontal
	NS = []int{-1, 0, 1, 0} // vertical

	NE = []int{-1, 0, 0, 1} // north<->east
	SE = []int{1, 0, 0, 1}  // south<->east

	NW = []int{-1, 0, 0, -1} // north<->west
	SW = []int{1, 0, 0, -1}  // south<->west

	NONE = []int{}

	T = map[rune][]int{
		'|': NS,
		'-': WE,
		'L': NE,
		'J': NW,
		'7': SW,
		'F': SE,
		'.': NONE,
	}
)

type coord struct {
	y, x int
}

type tubemap struct {
	m     map[coord]rune
	start coord
	w, h  int
}

func parse(input []string) tubemap {
	var m tubemap
	m.h = len(input)
	m.m = make(map[coord]rune)

	for y, line := range input {
		m.w = len(line)
		for x, c := range line {
			if c == 'S' {
				m.start = coord{y, x}
			}
			m.m[coord{y, x}] = c
		}
	}
	return m
}

func valid(t *tubemap, from, to coord) bool {
	toM, ok := T[t.m[to]]

	if !ok || len(toM) == 0 {
		return false
	}

	walkable := coord{y: to.y + toM[0], x: to.x + toM[1]}
	if walkable == from {
		return true
	}
	walkable = coord{y: to.y + toM[2], x: to.x + toM[3]}
	return walkable == from
}

func moves(t *tubemap, cl, cr coord, visited map[coord]rune) (left coord, right coord) {
	if cl == t.start {
		hasLeft := false
		var NM = []int{-1, 0, 0, -1, 0, 1, 1, 0}

		for ; len(NM) > 0; NM = NM[2:] {
			mm := coord{y: cl.y + NM[0], x: cl.x + NM[1]}
			if valid(t, cl, mm) && !hasLeft {
				left = mm
				hasLeft = true
			} else if valid(t, cl, mm) && hasLeft {
				right = mm
				break
			}
		}
		return left, right
	} else {
		nl := T[t.m[cl]]

		if len(nl) == 0 {
			fmt.Println(cl)
			fmt.Println(string(t.m[cl]))
		}
		// fmt.Printf("%v\n", cl)
		left = coord{y: cl.y + nl[0], x: cl.x + nl[1]}
		if _, ok := visited[left]; ok {
			left = coord{y: cl.y + nl[2], x: cl.x + nl[3]}
		}

		nr := T[t.m[cr]]
		right = coord{y: cr.y + nr[0], x: cr.x + nr[1]}
		if _, ok := visited[right]; ok {
			right = coord{y: cr.y + nr[2], x: cr.x + nr[3]}
		}

		return left, right
	}
}

func (t *tubemap) silver() int {
	visited := make(map[coord]rune)
	currentL, currentR := t.start, t.start
	c := 0
	for {
		visited[currentL] = t.m[currentL]
		visited[currentR] = t.m[currentL]
		c++
		left, right := moves(t, currentL, currentR, visited)

		if left == right {
			return c
		} else {
			currentL = left
			currentR = right
		}
	}
}

func (t *tubemap) gold() int {
	visited := make(map[coord]rune)
	currentL, currentR := t.start, t.start
	for {
		visited[currentL] = t.m[currentL]
		visited[currentR] = t.m[currentR]
		left, right := moves(t, currentL, currentR, visited)

		if left == right {
			visited[left] = t.m[left]
			visited[right] = t.m[right]
			break
		} else {
			currentL = left
			currentR = right
		}
	}

	score := 0
	m := make(map[coord]string)
	for y := 0; y < t.h; y++ {
		cross := 0
		for x := 0; x < t.w; x++ {
			c := coord{y, x}
			m[c] = string(t.m[c])
			if v, ok := visited[c]; ok {
				switch v {
				case '|':
					cross++
				case '7':
					cross++
				case 'F':
					cross++
				default:

				}
			} else {
				if cross%2 == 1 {
					score++
					m[c] = "I"
				} else {
					m[c] = "O"
				}
			}
		}
	}
	return score
}

func silver(input []string) int {
	tm := parse(input)
	return tm.silver()
}

func gold(input []string) int {
	tm := parse(input)
	return tm.gold()
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(GoldTestInput3))
	fmt.Printf("SILVER: %d\n", silver(utils.Input("day10/input")))
	fmt.Printf("TEST GOLD: %d\n", gold(GoldTestInput1))
	fmt.Printf("TEST GOLD: %d\n", gold(GoldTestInput2))
	fmt.Printf("TEST GOLD: %d\n", gold(GoldTestInput3))
	fmt.Printf("TEST GOLD: %d\n", gold(GoldTestInput4))
	fmt.Printf("GOLD: %d\n", gold(utils.Input("day10/input")))

}
