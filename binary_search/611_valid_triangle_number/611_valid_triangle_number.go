package validtrianglenumber

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			left := j + 1
			right := n - 1
			for left <= right {
				mid := (left + right) / 2
				if nums[mid] >= nums[i]+nums[j] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}

			// if left == n {
			//     continue
			// }

			count += left - j - 1

		}
	}

	return count

}
