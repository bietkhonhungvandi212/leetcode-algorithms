package rangesumquerymutable

type NumArray struct {
	SegmentTreeArr []int
	Length         int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	if n == 0 {
		return NumArray{SegmentTreeArr: []int{}, Length: 0}
	}

	arr := make([]int, n*4)
	segmentArr := make([]int, n*4)
	build(&segmentArr, 1, 0, n-1, nums)
	build(&arr, 1, 0, n-1, nums)
	tmp := NumArray{segmentArr, n}

	return tmp
}

func build(segmentArr *[]int, i, left, right int, nums []int) int {
	if left > right {
		return 0
	}

	if left == right {
		if i >= 0 && i < len(*segmentArr) && left >= 0 && left < len(nums) {
			(*segmentArr)[i] = nums[left]
			return nums[left]
		} else {
			return 0
		}
	}

	mid := left + (right-left)/2
	tleft := build(segmentArr, 2*i, left, mid, nums)
	tright := build(segmentArr, 2*i+1, mid+1, right, nums)

	if i >= 0 && i < len(*segmentArr) {
		(*segmentArr)[i] = tleft + tright
		return (*segmentArr)[i]
	} else {
		return 0
	}
}

func (this *NumArray) Update(index int, val int) {
	if index >= this.Length {
		return
	}

	var updateTreeArr func(indexTree, left, right, index, val int)
	updateTreeArr = func(indexTree, left, right, index, val int) {
		if left > right {
			return
		}

		if left == right {
			this.SegmentTreeArr[indexTree] = val
			return
		}

		mid := (left + right) / 2
		if index <= mid {
			updateTreeArr(2*indexTree, left, mid, index, val)
		} else {
			updateTreeArr(2*indexTree+1, mid+1, right, index, val)
		}

		this.SegmentTreeArr[indexTree] = this.SegmentTreeArr[2*indexTree] + this.SegmentTreeArr[2*indexTree+1]
	}

	updateTreeArr(1, 0, this.Length-1, index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > right {
		return 0
	}

	var min func(num1, num2 int) int
	var max func(num1, num2 int) int

	min = func(num1, num2 int) int {
		if num1 < num2 {
			return num1
		}

		return num2
	}

	max = func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}

		return num2
	}

	var getSumRange func(indexTree, tLeft, tRight, left, right int) int
	getSumRange = func(indexTree, tLeft, tRight, left, right int) int {
		if tLeft > tRight {
			return 0
		}

		if left > tRight || right < tLeft {
			return 0
		}

		if tLeft >= left && tRight <= right {
			return this.SegmentTreeArr[indexTree]
		}

		mid := (tLeft + tRight) / 2

		return getSumRange(2*indexTree, tLeft, mid, left, min(mid, right)) + getSumRange(2*indexTree+1, mid+1, tRight, max(mid+1, left), right)
	}

	return getSumRange(1, 0, this.Length-1, left, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
