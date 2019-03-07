package main

import "fmt"

func main() {
	words := []string{
		"flower",
		"flow",
		"flight",
	}
	fmt.Println(longestCommonPrefix(words))
}
