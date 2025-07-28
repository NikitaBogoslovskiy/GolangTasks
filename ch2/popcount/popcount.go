package popcount

var pc [256]uint8

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + uint8(i&1)
	}
}

func PopCount1(x uint64) int {
	return int(pc[uint8(x>>(0*8))] +
		pc[uint8(x>>(1*8))] +
		pc[uint8(x>>(2*8))] +
		pc[uint8(x>>(3*8))] +
		pc[uint8(x>>(4*8))] +
		pc[uint8(x>>(5*8))] +
		pc[uint8(x>>(6*8))] +
		pc[uint8(x>>(7*8))])
}

func PopCount2(x uint64) (res int) {
	for i := 0; i < 8; i++ {
		res += int(pc[uint8(x>>(i*8))])
	}
	return
}

func PopCount3(x uint64) (res int) {
	for i := 0; i < 64; i++ {
		res += int((x >> i) & 1)
	}
	return
}
