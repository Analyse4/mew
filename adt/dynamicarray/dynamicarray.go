package dynamicarray

import (
	"log"
)

type dynamicarray struct {
	buffer [1]interface{}
	len    int
	cap    int
}

func (da *dynamicarray) get(i int) interface{} {
	if i >= da.len {
		log.Fatalf("out of index: len: %d, index: %d\n", da.len, i)
	}
	return da.buffer[i]
}

func (da *dynamicarray) set(i int, v interface{}) {
	if i >= da.cap {
		da.resize(2 * da.cap)
	}
	da.buffer[i] = v
}

// TODO: golang is static language, it seems not capable of declare array in runtime
// length is part of array type, it must constant
func (da *dynamicarray) resize(max int) {

}
