package main

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}
