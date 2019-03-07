package main

func numJewelsInStones(J string, S string) int {
	counter := 0
	for i := 0; i < len(J); i++ {
		for j := 0; j < len(S); j++ {
			if J[i] == S[j] {
				counter++
			}
		}
	}
	return counter
}
