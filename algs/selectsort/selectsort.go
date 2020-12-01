package selectsort

// O(N^2)
func SelectSort(l []int) []int {
	for i := 0; i < len(l)-1; i++ {
		min := i
		for j := i; j < len(l); j++ {
			if l[j] < l[min] {
				min = j
			}
		}
		l[i], l[min] = l[min], l[i]
	}
	return l
}
