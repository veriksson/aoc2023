package utils

import (
	"os"
	"strings"
)

func Input(day string) []string {
	f, err := os.ReadFile("./" + day)
	if err != nil {
		panic(err)
	}

	sf := string(f)
	lines := strings.Split(sf, "\n")

	return lines[:len(lines)-1]
}
