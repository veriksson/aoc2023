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
	for _, num := range SplitTrim(line, " ") {
		ret = append(ret, Atoi(num))
	}
	return ret
}
