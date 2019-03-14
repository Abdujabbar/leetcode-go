package main

func checkPerfectNumber(num int) bool {
	perfectNumbers := map[int]bool{
		6:          true,
		28:         true,
		496:        true,
		8128:       true,
		33550336:   true,
		8589869056: true,
	}
	_, ok := perfectNumbers[num]
	return ok
}
