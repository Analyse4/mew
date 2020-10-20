package symboltable

import (
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
	h.Write(reflect.ValueOf(k).Bytes())
	index := h.Sum64() % maxbuffer
	if st.buffer[index] == nil {
		st.buffer[index] = new(linkedlist.Linkedlist)
	}
	st.buffer[index].Insert(k, v)
	h.Reset()
}

func (st *symboltable) Get(k interface{}) interface{} {
	if reflect.TypeOf(k) != st.key {
		log.Fatalf("type mismatch, want: %v, get: %v\n", st.key, reflect.ValueOf(k).Elem())
	}
	h.Write(reflect.ValueOf(k).Bytes())
	index := h.Sum64() % maxbuffer
	if st.buffer[index] == nil {
		h.Reset()
		return nil
	}
	v := st.buffer[index].Get(k)
	h.Reset()
	return v
}

func (st *symboltable) Delete(k interface{}) {
	if reflect.TypeOf(k) != st.key {
		log.Fatalf("type mismatch, want: %v, get: %v\n", st.key, reflect.ValueOf(k).Elem())
	}
	index := h.Sum64() % maxbuffer
	if st.buffer[index] == nil {
		h.Reset()
		return
	}
	st.buffer[index].Delete(k)
	h.Reset()
}

func (st *symboltable) IsEmpty() bool {
	return st.Size() == 0
}

func (st *symboltable) Contains(k interface{}) bool {
	if reflect.TypeOf(k) != st.key {
		log.Fatalf("type mismatch, want: %v, get: %v\n", st.key, reflect.ValueOf(k).Elem())
	}
	h.Write(reflect.ValueOf(k).Bytes())
	index := h.Sum64() % maxbuffer
	if st.buffer[index] == nil {
		h.Reset()
		return false
	}
	h.Reset()
	return nil != st.buffer[index].Get(k)
}

func (st *symboltable) Size() int {
	var len int
	for _, q := range st.buffer {
		len += q.Size()
	}
	return len
}
