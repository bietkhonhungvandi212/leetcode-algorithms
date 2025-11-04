package findxsumofallklongsubarraysi

import "sort"

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	freq := make(map[int]int)

	for i := 0; i < k; i++ {
		freq[nums[i]]++
	}
	left, right := 0, k-1
	for left < n-k+1 {
		key := nums[left]
		nums[left] = calXSum(freq, x)
		freq[key]--
		if freq[key] == 0 {
			delete(freq, key)
		}
		right++
		left++
		if right < n {
			freq[nums[right]]++

		}
	}

	return nums[0 : n-k+1]
}

func calXSum(freq map[int]int, x int) int {
	arr := make([][2]int, len(freq))

	for k, v := range freq {
		arr = append(arr, [2]int{k, v})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] > arr[j][0]
		}

		return arr[i][1] > arr[j][1]
	})

	sum := 0
	for i := range min(len(freq), x) {
		sum += arr[i][0] * arr[i][1]
	}

	return sum
}
