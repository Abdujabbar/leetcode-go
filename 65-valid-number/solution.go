package main

import (
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if s[0] == '.' {
		s = "0" + s
	}
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}
