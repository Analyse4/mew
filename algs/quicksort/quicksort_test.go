package quicksort

import (
	"reflect"
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

func TestQuickSortV2(t *testing.T) {
	type args struct {
		l *[]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"2-test", args{&[]int{5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5}},
		{"2-test", args{&[]int{5, 7, 6, 9, 0, 1, 2, 4, 3, 8}}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSortV2(tt.args.l)
			if !reflect.DeepEqual(*tt.args.l, tt.want) {
				t.Fatalf("not equal, got: %v, want: %v\n", *tt.args.l, tt.want)
			}
		})
	}
}
