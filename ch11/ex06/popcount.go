package popcount

// pc[i]はiのポピュレーションカウントです．
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountはxのポピュレーションカウント（1が設定されているビット数）を返します．
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount64(x uint64) int {
	var result int
	for i := 0; i < 64; i++ {
		result += int(byte((x >> i) & 1))
	}
	return result
}

func ClearPopCount(x uint64) int {
	var result int
	for result < 64 {
		if x == 0 {
			break
		}
		x = x & (x - 1)
		result++
	}
	return result
}
