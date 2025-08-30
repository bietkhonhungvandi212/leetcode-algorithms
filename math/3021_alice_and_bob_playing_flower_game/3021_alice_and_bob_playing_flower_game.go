package aliceandbobplayingflowergame

func flowerGame(n int, m int) int64 {
	xOdd := (n + 1) / 2
	xEven := n / 2
	yOdd := (m + 1) / 2
	yEven := m / 2

	return int64((xOdd * yEven) + (xEven * yOdd))
}
