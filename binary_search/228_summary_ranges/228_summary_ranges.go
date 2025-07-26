package summaryranges

import "strconv"

func summaryRanges(nums []int) []string {
	ranges := []string{}
	n := len(nums)
	start := 0
	end := 0

	for end < n-1 {
		if nums[end+1] == nums[end]+1 {
			end++
		} else {
			if start-end == 0 {
				ranges = append(ranges, strconv.Itoa(nums[start]))
			} else {
				ranges = append(ranges, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
			}
			end++
			start = end
		}
	}

	if start-end == 0 {
		ranges = append(ranges, strconv.Itoa(nums[start]))
	} else {
		ranges = append(ranges, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
	}

	return ranges
}
