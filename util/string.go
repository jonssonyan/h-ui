package util

import "strings"

func CompareVersion(version1, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	// Compare major version numbers
	if v1[0] > v2[0] {
		return 1
	} else if v1[0] < v2[0] {
		return -1
	}

	// If the major version numbers are the same, compare the minor version numbers
	if len(v1) > 1 && len(v2) > 1 {
		if v1[1] > v2[1] {
			return 1
		} else if v1[1] < v2[1] {
			return -1
		}
	}

	// If the major and minor versions are the same, compare the revision numbers
	if len(v1) > 2 && len(v2) > 2 {
		if v1[2] > v2[2] {
			return 1
		} else if v1[2] < v2[2] {
			return -1
		}
	}

	// The version number is exactly the same
	return 0
}
