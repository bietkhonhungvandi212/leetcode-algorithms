package rangesumqueryimmutable

type NumArray struct {
	segmentArr []int
	length     int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{segmentArr: []int{}, length: 0}
	}
	segmentArr := make([]int, 4*len(nums))
	n := len(nums)
	numArr := NumArray{segmentArr: segmentArr, length: n}
	numArr.build(1, 0, n-1, nums)
	return numArr
}

func (this *NumArray) build(v, left, right int, nums []int) int {
	if left == right {
		if v < len(this.segmentArr) && left < len(nums) {
			this.segmentArr[v] = nums[left]
			return this.segmentArr[v]
		}
		return 0
	}

	mid := (left + right) / 2
	// Bounds check for v before assigning
	if v < len(this.segmentArr) {
		sumLeft := this.build(2*v, left, mid, nums)
		sumRight := this.build(2*v+1, mid+1, right, nums)
		this.segmentArr[v] = sumLeft + sumRight
		return this.segmentArr[v]
	}
	return 0
}

func (this *NumArray) SumRange(left int, right int) int {
	if len(this.segmentArr) == 0 {
		return 0
	}

	return this.sumQuery(1, 0, this.length-1, left, right)
}

func (this *NumArray) sumQuery(v, tLeft, tRight, left, right int) int {
	if left > right {
		return 0
	}

	if tLeft == left && tRight == right {
		return this.segmentArr[v]
	}

	var min func(num1, num2 int) int
	var max func(num1, num2 int) int
	min = func(num1, num2 int) int {
		if num1 > num2 {
			return num2
		}

		return num1
	}

	max = func(num1, num2 int) int {
		if num1 < num2 {
			return num2
		}

		return num1
	}

	mid := (tLeft + tRight) / 2
	sumLeft := this.sumQuery(2*v, tLeft, mid, left, min(right, mid))
	sumRight := this.sumQuery(2*v+1, mid+1, tRight, max(mid+1, left), right)

	return sumLeft + sumRight
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
