package main

func getRow(rowIndex int) []int {
	result := []int{}

	result = append(result, 1)
	if rowIndex == 0 {
		return result
	}
	result = append(result, 1)
	if rowIndex == 1 {
		return result
	}

	rowIndex--
	sz := 1
	for rowIndex >= 0 {
		rowResult := []int{1}
		for i := 1; i < sz; i++ {
			rowResult = append(rowResult, result[i-1]+result[i])
		}
		rowResult = append(rowResult, 1)
		result = rowResult
		rowIndex--
		sz++
	}
	return result
}
