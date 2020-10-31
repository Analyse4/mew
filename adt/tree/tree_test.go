package tree

import "testing"

func TestHeight(t *testing.T) {
	tEX1 := New()
	tEX1.Insert(3)
	tEX1.Insert(2)
	tEX1.Insert(4)
	tEX1.Insert(1)
	tEX1.Insert(5)
	ra1 := 3

	ma1 := tEX1.Height()
	if ra1 != ma1 {
		t.Errorf("wrong height, want: %v, get: %v\n", ra1, ma1)
	}

	tEX2 := New()
	tEX2.Insert(1)
	tEX2.Insert(2)
	tEX2.Insert(3)
	tEX2.Insert(4)
	tEX2.Insert(5)
	ra2 := 5

	ma2 := tEX2.Height()
	if ra2 != ma2 {
		t.Errorf("wrong height, want: %v, get: %v\n", ra2, ma2)
	}
}
