package util

import "strings"

func CompareVersion(version1, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	// 比较主版本号
	if v1[0] > v2[0] {
		return 1
	} else if v1[0] < v2[0] {
		return -1
	}

	// 如果主版本号相同，比较次版本号
	if len(v1) > 1 && len(v2) > 1 {
		if v1[1] > v2[1] {
			return 1
		} else if v1[1] < v2[1] {
			return -1
		}
	}

	// 如果主版本号和次版本号都相同，则比较修订版本号
	if len(v1) > 2 && len(v2) > 2 {
		if v1[2] > v2[2] {
			return 1
		} else if v1[2] < v2[2] {
			return -1
		}
	}

	// 版本号完全相同
	return 0
}
