package main

func isValid(s string) bool {
	hmap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	opens := []byte{}
	for i := 0; i < len(s); i++ {
		if _, ok := hmap[s[i]]; ok {
			opens = append(opens, s[i])
		} else {
			if len(opens) == 0 {
				return false
			}
			if hmap[opens[len(opens)-1]] != s[i] {
				return false
			}
			opens = opens[:len(opens)-1]
		}
	}
	return len(opens) == 0
}
