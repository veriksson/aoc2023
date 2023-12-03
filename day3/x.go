package main

import "fmt"
import "unicode"
import "strings"
import "aoc2023/utils"

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
	w, h int
}

func parse(input []string) *board {
	var vals []rune
	var syms []int
	var w int = len(input[0])
	var h int = len(input)

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			chr := rune(input[y][x])
			vals = append(vals, chr)
			if chr != '.' && !unicode.IsDigit(chr) {
				syms = append(syms, utils.XYToI(x, y, w))
			}
		}
	}

	return &board{
		vals,
		syms,
		w,
		h,
	}
}

var nm = [][]int{
	[]int{-1, -1},
	[]int{-1, 0},
	[]int{-1, 1},

	[]int{0, -1},
	[]int{0, 1},

	[]int{1, -1},
	[]int{1, 0},
	[]int{1, 1},
}

func neighbourhood(sym int, b *board) []int {
	var nums []int
	sx := utils.IToX(sym, b.w)
	sy := utils.IToY(sym, b.w)

	taken := make(map[int]struct{})

	for i := range nm {
		nx := sx + nm[i][1]
		ny := sy + nm[i][0]

		if nx < 0 || nx > b.w {
			continue
		}

		if ny < 0 || ny > b.h {
			continue
		}

		n := utils.XYToI(nx, ny, b.w)
		if _, ok := taken[n]; ok {
			continue
		}

		v := b.vals[n]
		if v == '.' {
			continue
		}
		num, start, end := expand(nx, ny, b)

		nums = append(nums, num)
		for ; start <= end; start++ {
			ts := utils.XYToI(start, ny, b.w)
			taken[ts] = struct{}{}
		}
	}

	return nums
}

func expand(x, y int, b *board) (int, int, int) {
	var sb strings.Builder
	start := x
	end := x
	for {
		if start == 0 {
			break
		}

		n := utils.XYToI(start-1, y, b.w)
		if unicode.IsDigit(b.vals[n]) {
			start--
		} else {
			break
		}
	}

	for {
		if end == b.w-1 {
			break
		}

		n := utils.XYToI(end+1, y, b.w)
		if unicode.IsDigit(b.vals[n]) {
			end++
		} else {
			break
		}
	}

	orig := start
	for ; start <= end; start++ {
		n := utils.XYToI(start, y, b.w)
		sb.WriteRune(b.vals[n])
	}

	vv := string(b.vals[utils.XYToI(x, y, b.w)])
	return utils.Atoi(sb.String()), orig, end
}

func silver(input []string) int {
	b := parse(input)
	sum := 0
	for _, sym := range b.syms {
		ns := neighbourhood(sym, b)
		for _, n := range ns {
			sum += n
		}
	}
	return sum
}

func main() {
	fmt.Printf("TEST SILVER: %d\n", silver(TestInput))
}
