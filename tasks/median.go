package tasks

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := append(nums1, nums2...)
	sort.Ints(arr)

	if len(arr)%2 == 0 {
		return float64(arr[len(arr)/2-1]+arr[len(arr)/2]) / 2
	}
	return float64(arr[len(arr)/2])
}
