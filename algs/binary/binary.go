package binary

func GetBit(num, i int) bool {
	return num&(1<<(i-1)) > 0
}

func SetBit(num, i int) int {
	return num | 1<<(i-1)
}

func ClearBit(num, i int) int {
	return num & ^(1 << (i - 1))
}

func ClearBitsMSBThoughI(num, i int) int {
	return num&1<<(i-1) - 1
}

func ClearBitsIThrough0(num, i int) int {
	return num & (-1 << i)
}

func UpdateBit(num, i int, bitIs1 bool) int {
	tmp := ClearBit(num, i)
	if bitIs1 {
		tmp |= 1 << (i - 1)
	}
	return tmp
}
