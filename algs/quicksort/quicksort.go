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

// O(N^2)
func QuickSortV2(l *[]int) {
	if len(*l) <= 1 {
		return
	}
	sp := splitV2(l)
	tmpL := (*l)[:sp]
	QuickSortV2(&tmpL)
	if sp+1 < len(*l) {
		tmpR := (*l)[sp+1:]
		QuickSortV2(&tmpR)
	}
}

func splitV2(l *[]int) int {
	sp := 0
	start := 1
	end := len(*l) - 1
	for {
		for start <= end {
			if (*l)[start] > (*l)[sp] {
				break
			}
			start++
		}
		for end > start {
			if (*l)[end] < (*l)[sp] {
				break
			}
			end--
		}
		if end < start {
			break
		}
		(*l)[start], (*l)[end] = (*l)[end], (*l)[start]
		start++
		end--
	}
	(*l)[sp], (*l)[end] = (*l)[end], (*l)[sp]
	return end
}
