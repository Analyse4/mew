package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1-test", args{[]int{5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5}},
		{"2-test", args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
