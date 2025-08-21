package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	endIndex1 := m - 1
	endIndex2 := n - 1
	currentIndex := m + n - 1

	for endIndex1 >= 0 && endIndex2 >= 0 {
		if nums1[endIndex1] > nums2[endIndex2] {
			nums1[currentIndex] = nums1[endIndex1]
			endIndex1--
		} else {
			nums1[currentIndex] = nums2[endIndex2]
			endIndex2--
		}
		currentIndex--
	}

	for endIndex1 >= 0 {
		nums1[currentIndex] = nums1[endIndex1]
		endIndex1--
		currentIndex--
	}

	for endIndex2 >= 0 {
		nums1[currentIndex] = nums2[endIndex2]
		endIndex2--
		currentIndex--
	}
}
