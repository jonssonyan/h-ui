package util

func ArrContain[T comparable](arr []T, key T) bool {
	for _, item := range arr {
		if item == key {
			return true
		}
	}
	return false
}
