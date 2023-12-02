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
	items := strings.Split(s, sep)
	for i := range items {
		items[i] = strings.TrimSpace(items[i])
	}
	return items
}