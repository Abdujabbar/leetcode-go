package main

func isSubstractable(s string, index int) bool {
	return (s[index] == 'I' && (s[index+1] == 'V' || s[index+1] == 'X') ||
		(s[index] == 'X' && (s[index+1] == 'L' || s[index+1] == 'C')) ||
		(s[index] == 'C' && (s[index+1] == 'D' || s[index+1] == 'M')))
}

func romanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	index := 0
	for {
		if index >= len(s) {
			break
		}

		if index+1 < len(s) && isSubstractable(s, index) {
			result += symbols[s[index+1]] - symbols[s[index]]
			index += 2
			continue
		}

		result += symbols[s[index]]
		index++
	}
	return result
}
