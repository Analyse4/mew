package quicksort

import (
	"sort"
	"testing"
)

func TestQuickTest(t *testing.T) {
	flag := true
	e := []int{5, 7, 6, 9, 0, 1, 2, 4, 3, 8}
	ra := make([]int, len(e))
	ma := make([]int, len(e))
	copy(ra, e)
	copy(ma, e)
	//sort.Slice(ra, func(i, j int) bool { return i < j })
	sort.Ints(ra)
	QuickSort(ma)
	for i := range ra {
		if ra[i] != ma[i] {
			flag = false
		}
	}
	if !flag {
		t.Errorf("sort wrong, want: %v, get: %v\n", ra, ma)
	}
}
