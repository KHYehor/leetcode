package tasks

// Total complexity 2*O(n) + K
func topKFrequent(nums []int, k int) (result []int) {
	// Load the data to collect all cases
	// O(n)
	table := make(map[int]int)
	for _, key := range nums {
		if value, exist := table[key]; exist {
			table[key] = value + 1
		} else {
			table[key] = 1
		}
	}

	// Build buckets O(n)
	buckets := make([][]int, len(nums)+1)
	for key, v := range table {
		buckets[v] = append(buckets[v], key)
	}
	// Collect top K
	result = make([]int, 0, k)
	for i := len(buckets) - 1; i > 0; i-- {
		for _, v := range buckets[i] {
			result = append(result, v)
			if len(result) == k {
				return
			}
		}
	}

	return
}
