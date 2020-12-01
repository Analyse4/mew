package selectsort

import (
	"reflect"
	"testing"
)

func TestSelectSort(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1-test", args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectSort(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
