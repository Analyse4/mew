package linkedlist

import "testing"

func TestQueue(t *testing.T) {
	ra := 0
	for i := 0; i < 100; i++ {
		ra += i
	}
	ma := 0

	ll := New()
	for i := 0; i < 100; i++ {
		ll.Insert(i, i)
	}
	for i := 0; i < 100; i++ {
		if ll.Size() == 0 {
			t.Error("size is 0, want: size is not 0, get: size is 0\n")
		}
		ma += ll.Get(i).(int)
	}
	if ma != ra {
		t.Errorf("result is wrong, want: %d, get: %d\n", ra, ma)
	}
}
