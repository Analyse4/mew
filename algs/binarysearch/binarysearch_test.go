package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	e1 := make([]int, 100)
	e2 := make([]int, 100)
	e3 := make([]int, 100)
	for i := 0; i < len(e1); i++ {
		e1[i] = i
		e2[i] = 1
		e3[i] = 1000
	}
	ra1 := 1

	i1 := BinarySearch(e1[ra1], e1)
	if i1 != ra1 {
		t.Errorf("index is wrong, want: %v, get: %v\n", ra1, i1)
	}
	i2 := BinarySearch(50, e2)
	if i2 > 0 {
		t.Errorf("index is wrong, want: %v, get: %v\n", -1, i2)
	}
	i3 := BinarySearch(50, e3)
	if i3 > 0 {
		t.Errorf("index is wrong, want: %v, get: %v\n", -1, i3)
	}
}
