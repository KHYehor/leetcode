package tasks

// https://leetcode.com/problems/task-scheduler/

func isBufferContains(buffer []byte, task byte) bool {
	for _, b := range buffer {
		if b == task {
			return true
		}
	}
	return false
}

func leastInterval(tasks []byte, n int) (res int) {
	m := map[byte]int{}
	// Make a map of tasks
	for _, task := range tasks {
		if _, ok := m[task]; !ok {
			m[task] = 0
		}
		m[task]++
	}

	// To store the moving buffer
	buffer := make([]byte, 0, n)

	for i := 0; i < len(tasks); i++ {
		valueMax := -1
		keyMax := byte(0)

		for key, value := range m {
			// Check next possible tasks
			if isBufferContains(buffer, key) {
				continue
			}

			// Use first the most often
			if value > valueMax {
				valueMax = value
				keyMax = key
			}
		}

		// Set idle operation
		if valueMax == -1 {
			res++
			buffer = append(buffer, byte(0))
			// Iterate again
			i--
		} else {
			// Add the task to the execution
			res++
			// Extend the buffer with the task
			buffer = append(buffer, keyMax)
			// If the value is fully used, delete it from the cache
			if m[keyMax]--; m[keyMax] == 0 {
				delete(m, keyMax)
			}
		}

		// Move the buffer
		if len(buffer) > n {
			buffer = buffer[1:]
		}
	}

	return
}
