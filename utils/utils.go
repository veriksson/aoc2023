package utils

import (
	"strconv"
	"strings"
)

func Atoi(s string) int {
	a, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return a
}

func SplitTrim(s, sep string) []string {
	var ret []string
	for _, item := range strings.Split(s, sep) {
		item = strings.TrimSpace(item)
		if item != "" {
			ret = append(ret, item)
		}
	}
	return ret
}

func IToX(i, W int) int {
	return i % W
}

func IToY(i, W int) int {
	return i / W
}

func XYToI(x, y, W int) int {
	return y*W + x
}

func IntsOfString(line string) []int {
	var ret []int
	for _, num := range strings.Fields(line) {
		ret = append(ret, Atoi(num))
	}
	return ret
}

// taken from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Single(input []string, i int) []string {
	return []string{input[i]}
}
