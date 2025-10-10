package takingmaximumenergyfromthemysticdungeon

import "math"

func maximumEnergy(energy []int, k int) int {
	//This process will be repeated until you reach the magician where (i + k) does not exist.
	maxGain := math.MinInt
	n := len(energy)

	for i := range n {
		if i-k >= 0 {
			energy[i] = max(energy[i]+energy[i-k], energy[i])
		}

		if i >= n-k {
			maxGain = max(maxGain, energy[i])
		}
	}

	return maxGain
}
