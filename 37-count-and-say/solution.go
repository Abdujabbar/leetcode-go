package main

import "strconv"

func countAndSay(n int) string {
	r := "1"
	j := 1
	for n > 1 {
		s := ""
		for i := 0; i < len(r); i++ {
			for j = i + 1; j < len(r); j++ {
				if r[i] != r[j] {
					break
				}
			}
			s += strconv.Itoa(j-i) + string(r[i])
			i = j - 1
		}
		r = s
		n--
	}
	return r
}
