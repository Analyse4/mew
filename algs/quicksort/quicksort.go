package quicksort

// TODO: generic
func QuickSort(list []int) {
	if len(list) <= 1 {
		return
	}
	sl := split(list)
	QuickSort(list[:sl])
	QuickSort(list[sl+1:])
}

func split(list []int) int {
	v := list[0]
	i := 1
	j := len(list) - 1
	for {
		for ; i < len(list); i++ {
			if list[i] > v {
				break
			}
		}
		if i >= len(list) {
			break
		}
		for ; j > 0; j-- {
			if list[j] < v {
				break
			}
		}
		if j < i {
			break
		}
		swap(i, j, list)
		i++
		j--
	}
	if j != 0 {
		swap(0, j, list)
	}
	return j
}

func swap(v1, v2 int, list []int) {
	tmp := list[v1]
	list[v1] = list[v2]
	list[v2] = tmp
}
