package heapM

type HeapM struct {
	h []int
}

// O(Nlog(N))
func New(l []int) *HeapM {
	heap := &HeapM{make([]int, 0)}
	for i := len(l) / 2; i > 0; i-- {
		min := 2 * i
		if 2*i+1 <= len(l) {
			if l[2*i-1] > l[2*i+1-1] {
				min = 2*i + 1
			}
		}
		if l[i-1] > l[min-1] {
			l = sink(l, i)
		}
	}
	heap.h = append(heap.h, l...)
	return heap
}

func (h *HeapM) Insert(num int) {
	h.h = append(h.h, num)
	h.h = swim(h.h)
}

// O(log(N))
func swim(l []int) []int {
	target := len(l) - 1
	for target/2 >= 0 && l[target] < l[target/2] {
		l[target], l[target/2] = l[target/2], l[target]
		target = target / 2
	}
	return l
}

func (h *HeapM) Min() int {
	min := h.h[0]
	h.h[0], h.h[len(h.h)-1] = h.h[len(h.h)-1], h.h[0]
	h.h = h.h[:len(h.h)-1]
	h.h = sink(h.h, 1)
	return min
}

// O(logN)
func sink(l []int, index int) []int {
	target := index
	for target*2 <= len(l) {
		tmpMin := 2 * target
		if 2*target+1 <= len(l) {
			if l[2*target-1] > l[2*target+1-1] {
				tmpMin = 2*target + 1
			}
		}
		if l[target-1] > l[tmpMin-1] {
			l[target-1], l[tmpMin-1] = l[tmpMin-1], l[target-1]
			target = tmpMin
		} else {
			return l
		}
	}
	return l
}
