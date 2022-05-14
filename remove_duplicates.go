package main

func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	end := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[end] {
			end++
			nums[end] = nums[i]
		}
	}
	return end + 1
}
