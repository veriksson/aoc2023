package main

import "fmt"

var TestInput = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

type board struct {
	vals []rune
	syms []int
}

func parse(input []string) *board {
	var vals []rune
	var syms []int

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			chr := rune(input[y][x])
			vals = append(vals, chr)
			if chr != '.' && !unicode.IsDigit(chr) {
				syms = append(syms, utils.XYToI(x, y, len(input[y])))
			}
		}
	}

	return &board{
		vals,
		syms,
	}
}

func silver(input []string) int {
	return 0
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput))
}
