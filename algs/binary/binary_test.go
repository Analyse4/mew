package binary

import "testing"

func TestGetBit(t *testing.T) {
	type args struct {
		num int
		i   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"5-3-test", args{5, 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBit(tt.args.num, tt.args.i); got != tt.want {
				t.Errorf("GetBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetBit(t *testing.T) {
	type args struct {
		num int
		i   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5-2-test", args{5, 2}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetBit(tt.args.num, tt.args.i); got != tt.want {
				t.Errorf("SetBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearBit(t *testing.T) {
	type args struct {
		num int
		i   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5-3-test", args{5, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearBit(tt.args.num, tt.args.i); got != tt.want {
				t.Errorf("ClearBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearBitsMSBThoughI(t *testing.T) {
	type args struct {
		num int
		i   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5-2-test", args{5, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearBitsMSBThoughI(tt.args.num, tt.args.i); got != tt.want {
				t.Errorf("ClearBitsMSBThoughI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearBitsIThrough0(t *testing.T) {
	type args struct {
		num int
		i   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5-2-test", args{5, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearBitsIThrough0(tt.args.num, tt.args.i); got != tt.want {
				t.Errorf("ClearBitsIThrough0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateBit(t *testing.T) {
	type args struct {
		num    int
		i      int
		bitIs1 bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5-4-test", args{5, 4, true}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateBit(tt.args.num, tt.args.i, tt.args.bitIs1); got != tt.want {
				t.Errorf("UpdateBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
