package maximumvalueofanorderedtripleti

func maximumTripletValue(nums []int) int64 {
	var (
		maxTrplet int64 = 0
		maxDiff   int64 = 0
		maxEle    int64 = 0
	)

	for _, num := range nums {
		maxTrplet = MaxInt64(maxTrplet, maxDiff*int64(num))
		maxDiff = MaxInt64(maxDiff, maxEle-int64(num))
		maxEle = MaxInt64(maxEle, int64(num))

	}

	return maxTrplet
}

func maximumTripletValueOtherWay(nums []int) int64 {
	n := len(nums)
	maxEle := make([]int, n)
	maxMultiple := make([]int, n)

	for i := 1; i < n; i++ {
		maxEle[i] = MaxInt(nums[i], maxEle[i-1])
		maxMultiple[n-i-1] = MaxInt(nums[n-i], maxMultiple[n-i])
	}

	res := 0
	for i := 1; i < n-1; i++ {
		triplet := (maxEle[i] - nums[i]) * maxMultiple[i]
		res = MaxInt(res, triplet)
	}

	return int64(res)
}

func MaxInt64(num1 int64, num2 int64) int64 {
	if num1 > num2 {
		return num1
	}
	return num2
}

func MaxInt(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func main() {
	nums := []int{12, 6, 1, 2, 7}

	result := maximumTripletValue(nums)

	println(result)
}
