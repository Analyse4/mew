package mergesort

// Time: O(N^2) Space: O(N)
func MergeSort(l []int) []int {
	if len(l) == 1 {
		return l
	}
	middle := len(l) / 2
	left := MergeSort(l[:middle])
	right := MergeSort(l[middle:])
	return merge(left, right)
}

func merge(l, r []int) []int {
	target := make([]int, 0)
	left := 0
	right := 0
	for {
		if l[left] < r[right] {
			target = append(target, l[left])
			left++
		} else {
			target = append(target, r[right])
			right++
		}
		if left == len(l) {
			for i := right; i < len(r); i++ {
				target = append(target, r[i])
			}
			return target
		}
		if right == len(r) {
			for i := left; i < len(l); i++ {
				target = append(target, l[i])
			}
			return target
		}
	}
}
