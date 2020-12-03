package heapsort

import "github.com/Analyse4/mew/adt/heapM"

func HeapSort(l []int) []int {
	hp := heapM.New(l)
	sortedL := make([]int, 0)
	for i := 0; i < len(l); i++ {
		sortedL = append(sortedL, hp.Min())
	}
	return sortedL
}
