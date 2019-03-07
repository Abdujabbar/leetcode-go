package main

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j > len(haystack) {
				return -1
			}
			if haystack[i+j] != needle[j] {
				break
			}

		}
	}
	return -1
}
