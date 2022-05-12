package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, len(input))
	result = make([]string, len(input))
	i := 0
	for mapKey := range input {
		keys[i] = mapKey
		i++
	}
	sort.Ints(keys)
	for i := range keys {
		result[i] = input[keys[i]]
	}
	return
}
