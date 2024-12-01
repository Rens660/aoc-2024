package utils

import "sort"

func SortAsc(sl *[]int) {
	slice := *sl
	sort.Ints(slice)
}
