package main

func commonPrefix(s1, s2 string) string {
	index := 0
	for index < len(s1) && index < len(s2) {
		if s1[index] == s2[index] {
			index++
		} else {
			break
		}
	}
	return s1[0:index]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	if prefix != "" {
		for i := 1; i < len(strs); i++ {
			prefix = commonPrefix(prefix, strs[i])
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}
