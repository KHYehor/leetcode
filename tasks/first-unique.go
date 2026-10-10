package tasks

var textString = "testesabc"

func firstUnique(s string) int {
	// char -> how many times met
	var counts = make(map[rune]int)
	for _, v := range s {
		counts[v]++
	}
	// second pass in string order: the first char met once is the answer
	for i, v := range s {
		if counts[v] == 1 {
			return i
		}
	}

	return -1
}
