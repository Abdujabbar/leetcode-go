package main

//Kadane's algorithm
func maxSubArray(nums []int) int {
	maxEndingHere := 0
	maxSoFar := -2147483647
	for i := 0; i < len(nums); i++ {
		maxEndingHere += nums[i]
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar
}
