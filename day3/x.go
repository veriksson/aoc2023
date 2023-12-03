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

func silver(input []string) int {
	return 0
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput))
}
