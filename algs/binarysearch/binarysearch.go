package binarysearch

// TODO: generic
func BinarySearch(v int, list []int) int {
	midIndex := len(list) / 2
	if v == list[midIndex] {
		return midIndex
	} else if len(list) == 1 {
		return -1
	} else if v > list[midIndex] {
		if midIndex == len(list)-1 {
			return -1
		}
		return BinarySearch(v, list[midIndex+1:])
	} else {
		return BinarySearch(v, list[:midIndex])
	}
}
