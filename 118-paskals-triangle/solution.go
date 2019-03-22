package main

func generate(numRows int) [][]int {
	result := [][]int{}
	if numRows == 0 {
		return result
	}
	result = append(result, []int{1})
	if numRows == 1 {
		return result
	}
	result = append(result, []int{1, 1})
	if numRows == 2 {
		return result
	}
	numRows -= 2
	sz := 2
	for numRows > 0 {
		result = append(result, []int{})
		result[sz] = append(result[sz], 1)
		for i := 1; i < sz; i++ {
			result[sz] = append(result[sz], result[sz-1][i-1]+result[sz-1][i])
		}
		result[sz] = append(result[sz], 1)
		sz++
		numRows--
	}
	return result
}
