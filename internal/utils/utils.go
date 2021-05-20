package utils

func BatchCollection(collection []int, size int) [][]int {
	var batches [][]int
	if len(collection) == 0 || size <= 0 {
		return batches
	}

	index := 0
	for {
		remaining := len(collection) - index
		if remaining <= size {
			batches = append(batches, collection[index:index+remaining])
			break
		}
		batches = append(batches, collection[index:index+size])
		index += size
	}
	return batches
}

func SwapKeysAndValues(initial map[int]int) map[int]int {
	swapped := make(map[int]int, len(initial))
	for key, value := range initial {
		if _, found := swapped[value]; found {
			panic("Duplicate value detected. Make sure initial map has no duplicates.")
		}
		swapped[value] = key
	}
	return swapped
}

func Remove(collection []int, values ...int) []int {
	var cleared []int
	for _, value := range collection {
		if !contains(values, value) {
			cleared = append(cleared, value)
		}

	}
	return cleared
}

func contains(collection []int, value int) bool {
	for _, v := range collection {
		if v == value {
			return true
		}
	}
	return false
}
