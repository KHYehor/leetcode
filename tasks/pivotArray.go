package tasks

// Also workable
//func pivotArray(nums []int, pivot int) []int {
//	var less = make([]int, 0, len(nums)/3)
//	var great = make([]int, 0, len(nums)/3)
//	var equal = make([]int, 0, len(nums)/3)
//	for _, v := range nums {
//		if v < pivot {
//			less = append(less, v)
//		} else if v > pivot {
//			great = append(great, v)
//		} else {
//			equal = append(equal, v)
//		}
//	}
//	res := append(append(less, equal...), great...)
//	return res
//}

func pivotArray(nums []int, pivot int) []int {
	var res = make([]int, len(nums))
	var amount = len(nums) - 1
	var l = 0
	var g = len(nums) - 1
	for i, v := range nums {
		if v < pivot {
			res[l] = v
			l++
		}
		if r := nums[amount-i]; r > pivot {
			res[g] = r
			g--
		}
	}
	for l <= g {
		res[l] = pivot
		l++
	}

	return res
}
