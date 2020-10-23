// Package symboltalble implements a map backed by a hash table.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.

package symboltable

import (
	"bytes"
	"encoding/binary"
	"hash/maphash"
	"log"
	"reflect"

	"github.com/Analyse4/mew/adt/linkedlist"
)

const (
	maxbuffer = 1000
)

var String = reflect.TypeOf("")
var Int = reflect.TypeOf(0)

var h maphash.Hash

type symboltable struct {
	key    reflect.Type
	value  reflect.Type
	buffer []*linkedlist.Linkedlist
}

func New(k, v reflect.Type) *symboltable {
	return &symboltable{k, v, make([]*linkedlist.Linkedlist, maxbuffer)}
}

func (st *symboltable) Put(k, v interface{}) {
	if reflect.TypeOf(k) != st.key {
		log.Fatalf("type mismatch, want: %v, get: %v\n", st.key, reflect.ValueOf(k).Elem())
	}
	if reflect.TypeOf(v) != st.value {
		log.Fatalf("type mismatch, want: %v, get: %v\n", st.value, reflect.ValueOf(v).Elem())
	}
	index := keyToIndex(k)
	if st.buffer[index] == nil {
		st.buffer[index] = linkedlist.New()
	}
	st.buffer[index].Insert(k, v)
}

func (st *symboltable) Get(k interface{}) interface{} {
	if reflect.TypeOf(k) != st.key {
		log.Fatalf("type mismatch, want: %v, get: %v\n", st.key, reflect.ValueOf(k).Elem())
	}
	index := keyToIndex(k)
	if st.buffer[index] == nil {
		return nil
	}
	v := st.buffer[index].Get(k)
	return v
}

func (st *symboltable) Delete(k interface{}) {
	if reflect.TypeOf(k) != st.key {
		log.Fatalf("type mismatch, want: %v, get: %v\n", st.key, reflect.ValueOf(k).Elem())
	}
	index := keyToIndex(k)
	if st.buffer[index] == nil {
		return
	}
	st.buffer[index].Delete(k)
}

func (st *symboltable) IsEmpty() bool {
	return st.Size() == 0
}

func (st *symboltable) Contains(k interface{}) bool {
	if reflect.TypeOf(k) != st.key {
		log.Fatalf("type mismatch, want: %v, get: %v\n", st.key, reflect.ValueOf(k).Elem())
	}
	index := keyToIndex(k)
	if st.buffer[index] == nil {
		return false
	}
	return nil != st.buffer[index].Get(k)
}

func (st *symboltable) Size() int {
	var len int
	for _, q := range st.buffer {
		if q == nil {
			continue
		}
		len += q.Size()
	}
	return len
}

func keyToIndex(k interface{}) uint64 {
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, k)
	h.Write(buff.Bytes())
	index := h.Sum64() % maxbuffer
	h.Reset()
	return index
}
