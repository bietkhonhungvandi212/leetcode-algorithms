package solvingquestionswithbrainpower

func mostPoints(questions [][]int) int64 {
	mem := make([]int64, len(questions))

	return calculatePoints(questions, 0, mem)
}

func calculatePoints(questions [][]int, currentIndex int, mem []int64) int64 {
	if currentIndex > len(questions)-1 {
		return 0
	}

	if mem[currentIndex] != 0 {
		return mem[currentIndex]
	}

	process := int64(questions[currentIndex][0]) + calculatePoints(questions, currentIndex+questions[currentIndex][1]+1, mem)
	skipCurrent := calculatePoints(questions, currentIndex+1, mem)

	mem[currentIndex] = MaxInt(process, skipCurrent)

	return mem[currentIndex]
}

func MaxInt(num1 int64, num2 int64) int64 {
	if num1 > num2 {
		return num1
	}
	return num2
}

func main() {

}
