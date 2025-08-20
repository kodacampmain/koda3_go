package minitask1

func InsertIntoSlice[K string | bool | int](inputSlice []K, target, newValue K) []K {
	indexTarget := -1
	for i, v := range inputSlice {
		if v == target {
			indexTarget = i
			break
		}
	}
	if indexTarget == -1 {
		return inputSlice
	}

	results := make([]K, indexTarget+1)
	copy(results, inputSlice)
	results = append(results, newValue)
	results = append(results, inputSlice[indexTarget+1:]...)

	return results
}
