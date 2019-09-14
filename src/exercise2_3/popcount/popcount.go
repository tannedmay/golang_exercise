package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	count := 0
	var i uint64
	for ; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}
