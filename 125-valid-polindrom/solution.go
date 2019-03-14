package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return false
	}
	p := re.ReplaceAllString(s, "")
	reversed := ""
	for i := len(p) - 1; i >= 0; i-- {
		reversed += string(p[i])
	}
	return strings.ToLower(p) == strings.ToLower(reversed)
}
