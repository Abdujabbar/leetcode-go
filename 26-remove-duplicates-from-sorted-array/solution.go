package main

func removeDuplicates(nums []int) int {
	c := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[c] {
			c++
			nums[c] = nums[i]
		}
	}
	return c + 1
}
