package popcount

func PopCountWithRightBit(x uint64) int {
	count := 0
	for x > 0 {
		x &= x - 1
		count++
	}
	return count
}
