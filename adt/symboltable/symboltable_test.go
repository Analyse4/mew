package symboltable

import (
	"strconv"
	"testing"
)

func TestSymboltable(t *testing.T) {
	const loopcount = 100
	var ra string
	var ma string
	for i := 0; i < loopcount; i++ {
		ra += strconv.Itoa(i)
	}

	st := New(Int, String)
	for i := 0; i < loopcount; i++ {
		st.Put(i, strconv.Itoa(i))
	}
	if st.IsEmpty() {
		t.Errorf("API isEmpty is wrong, want: false, got: %v\n", st.IsEmpty())
	}
	if st.Size() != loopcount {
		t.Errorf("symboltable size is wrong, want: %v, got: %v\n", loopcount, st.Size())
	}
	for i := 0; i < 100; i++ {
		ma += st.Get(i).(string)
	}
	if ma != ra {
		t.Errorf("sum result is wrong, want: %v, got: %v\n", ra, ma)
	}
}
