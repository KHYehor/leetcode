package tasks

import "strconv"

func reverse(x int) int {
	// Check for 32+ bit int
	if x <= -2147483647 || x >= 2147483647 {
		return 0
	}

	str := strconv.Itoa(x)
	res := ""

	// Check for negative
	start := 0
	if str[0] == '-' {
		res += "-"
		start = 1
	}

	// Remove leading zeros
	for start < len(str) {
		if str[start] == '0' {
			start++
			continue
		}
		break
	}

	// Reverse
	for i := len(str) - 1; i >= start; i-- {
		res += string(str[i])
	}

	resInt, _ := strconv.Atoi(res)

	return resInt
}
