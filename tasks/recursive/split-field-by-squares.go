package recursive

func SplitFieldBySquares(width, height, size int) int {
	// Init first run
	if size == 0 {
		size = height
	}
	// Take mods of height and width
	modW := width % size
	// Keep on splitting by width
	if modW != 0 {
		return SplitFieldBySquares(width, height, modW)
	}
	// Keep on splitting by height
	modH := height % size
	if modH != 0 {
		return SplitFieldBySquares(width, height, modH)
	}

	// Ideal square form founded
	return size
}
