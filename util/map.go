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
			// 当当前组已满或达到最后一组时，将映射添加到结果切片中
			segments = append(segments, currentGroup)
			currentGroup = make(map[string]T) // 初始化一个新的映射
			groupIndex++
		}
	}

	return segments
}
