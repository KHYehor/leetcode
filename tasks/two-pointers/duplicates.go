package two_pointers

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=problem-list-v2&envId=two-pointers

func removeDuplicates(nums []int) int {
	// Empty array
	if len(nums) == 0 {
		return 0
	}
	// Only one value
	if len(nums) == 1 {
		return 1
	}
	// Define placeholder for unique elements
	res := make([]int, 0, len(nums))
	// Init first element
	prev := nums[0]
	res = append(res, prev)

	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			res = append(res, nums[i])
		}
		prev = nums[i]
	}
	fmt.Println("Before", nums)
	nums = res
	fmt.Println("After", nums)
	return len(nums)
}
