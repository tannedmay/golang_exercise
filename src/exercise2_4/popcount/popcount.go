package popcount

func PopCountWith64Shift(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x&1 > 0 {
			count += 1
		}
		x >>= 1
	}
	return count
}
