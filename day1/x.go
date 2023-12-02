package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"strings"
)

var TestInputSilver = []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
var TestInputGold = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func silver(line string) int {
	var sb strings.Builder
	for _, c := range line {
		if c >= '0' && c <= '9' {
			sb.WriteRune(c)
		}
	}
	full := sb.String()
	fulln := string(full[0]) + string(full[len(full)-1])
	num, err := strconv.Atoi(fulln)
	if err != nil {
		panic(err)
	}
	return num
}

func gold(line string) int {
	named := func(num string, i int) bool {
		l := len(num)
		if i+l > len(line) {
			return false
		}
		for j := 0; j < len(num); j++ {
			if line[i+j] != num[j] {
				return false
			}
		}
		return true
	}

	var sb strings.Builder
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c >= '0' && c <= '9' {
			sb.WriteByte(c)
			continue
		}

		if c != 'o' && c != 't' && c != 'f' && c != 's' && c != 'e' && c != 'n' {
			continue
		}

		if named("one", i) {
			sb.WriteRune('1')
			continue
		}

		if named("two", i) {
			sb.WriteRune('2')
			continue
		}

		if named("three", i) {
			sb.WriteRune('3')
			continue
		}

		if named("four", i) {
			sb.WriteRune('4')
			continue
		}
		if named("five", i) {
			sb.WriteRune('5')
			continue
		}
		if named("six", i) {
			sb.WriteRune('6')
			continue
		}
		if named("seven", i) {
			sb.WriteRune('7')
			continue
		}
		if named("eight", i) {
			sb.WriteRune('8')
			continue
		}
		if named("nine", i) {
			sb.WriteRune('9')
			continue
		}
	}

	full := sb.String()
	fulln := string(full[0]) + string(full[len(full)-1])
	num, err := strconv.Atoi(fulln)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	sums, sumg := 0, 0
	for _, l := range util.Input("day1/input") {
		sums += silver(l)
		sumg += gold(l)
	}
	fmt.Printf("SILVER:\t%d\n", sums)
	fmt.Printf("GOLD:\t%d\n", sumg)
}
