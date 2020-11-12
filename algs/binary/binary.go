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

// O(counts of target number's digits)
func GetCountsOfDigits(num int) int {
	if num == 0 {
		return 1
	}
	endNum := 1
	count := 0
	for endNum <= num {
		count++
		endNum <<= 1
	}
	return count
}
