package main

import (
	"aoc2023/utils"
	"fmt"
	"sort"
)

var TestInput = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func parse(input []string) [][]int {
	var ret [][]int
	for _, line := range input {
		ret = append(ret, utils.IntsOfString(line))
	}
	return ret
}

func differences(reading []int) [][]int {
	var ret [][]int
	check := true
	for check {
		var next []int
		for i := 1; i < len(reading); i++ {
			next = append(next, reading[i]-reading[i-1])
		}
		ret = append(ret, next)
		reading = next
		check = false
		for i := 0; i < len(reading); i++ {
			if reading[i] != 0 {
				check = true
				break
			}
		}
	}
	return ret
}

func predict(reading []int) int {
	diffs := differences(reading)
	sort.Slice(diffs, func(i, j int) bool { return i < j })
	ph := 0
	for _, diff := range diffs {
		ph = ph + diff[len(diff)-1]
	}
	return ph + reading[len(reading)-1]
}

var printDiffs = false

func predict2(reading []int) int {
	diffs := differences(reading)
	sort.Slice(diffs, func(i, j int) bool { return len(diffs[i]) < len(diffs[j]) })
	if printDiffs {
		for _, diff := range diffs {
			fmt.Printf("%v\n", diff)
		}
	}
	ph := 0
	for _, diff := range diffs {
		ph = -ph + diff[0]
	}
	return -ph + reading[0]
}

func silver(input []string) int {
	readings := parse(input)
	sum := 0
	for _, reading := range readings {
		sum += predict(reading)
	}
	return sum
}

func gold(input []string) int {
	readings := parse(input)
	sum := 0
	for _, reading := range readings {
		prediction := predict2(reading)
		fmt.Printf("prediction for %v: %d\n", reading, prediction)
		sum += prediction
	}
	return sum
}

func main() {
	// single := func(input []string, i int) []string {
	// 	return []string{input[i]}
	// }
	// fmt.Printf("TEST SILVER: %d\n", silver(TestInput))
	// fmt.Printf("SILVER: %d\n", silver(utils.Input("day9/input")))
	fmt.Printf("TEST GOLD: %d\n", gold(TestInput))
	fmt.Printf("GOLD: %d\n", gold(utils.Input("day9/input")))
}
