package tasks

func maxDistinct(s string) int {
	var m = make(map[int32]struct{}, len(s))
	var unique = 0
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			unique++
		}
	}
	return unique
}
