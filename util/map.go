package util

func SplitMap[T any](inputMap map[string]T, chunkSize int) []map[string]T {
	length := len(inputMap)
	quantity := (length + chunkSize - 1) / chunkSize
	segments := make([]map[string]T, 0, quantity)

	var groupIndex int
	currentGroup := make(map[string]T)

	for key, value := range inputMap {
		currentGroup[key] = value

		if len(currentGroup) == chunkSize || groupIndex+1 == quantity {
			// When the current group is full or the last group is reached, add the mapping to the result slice
			segments = append(segments, currentGroup)
			currentGroup = make(map[string]T) // Initialize a new mapping
			groupIndex++
		}
	}

	return segments
}
