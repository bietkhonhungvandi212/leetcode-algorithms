package compareversionnumbers

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	//The value of the revision is its integer conversion ignoring leading zeros.
	revision1s := strings.Split(version1, ".")
	revision2s := strings.Split(version2, ".")

	i := 1
	for i < max(len(revision1s), len(revision2s)) {
		var revis1 int
		var revis2 int
		if i < len(revision1s) {
			num, _ := strconv.Atoi(revision1s[i])
			revis1 = num
		}

		if i < len(revision2s) {
			num, _ := strconv.Atoi(revision2s[i])
			revis2 = num
		}

		if revis1 < revis2 {
			return -1
		}
		if revis1 > revis2 {
			return 1
		}
		i++
	}

	return 0
}
