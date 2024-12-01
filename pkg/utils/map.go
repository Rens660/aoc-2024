package utils

func MakeMap(nmbrs *[]int) map[int]int {
	m := make(map[int]int)

	for _, val := range *nmbrs {
		m[val] += 1
	}

	return m
}
