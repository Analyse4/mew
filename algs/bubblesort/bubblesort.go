package bubblesort

func BubbleSort(l []int) []int {
	for i := len(l) - 1; i > 0; i-- {
		for j := 0; j+1 <= i; j++ {
			if l[j] > l[j+1] {
				l[j+1], l[j] = l[j], l[j+1]
			}
		}
	}
	return l
}
