package findtheminimumamountoftimetobrewpotions

import "fmt"

func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	m := len(mana)
	// prefix[j] stores the earliest time wizard j completes their tasks after processing the current potion
	prefix := make([]int64, n)

	for i := range m {
		// use current as keeping track of the earliest time that a wizard can complete their task
		// Because we only brewing potion by pass it sequentially through each wizard
		// So current is as the earliest time that the all wizard can be completed brewing potion without being blocked by time to brew potion from previous potion
		current := int64(0)
		for j := range n {
			// wizard j can start brewing potion i after:
			// 1. The previous wizard (j-1) finishes
			// 2. Wizard j finishes their previous tasks (prefix[j])
			// Take the maximum to ensure the wizard is free, then add brewing time
			current = max(current+int64(skill[j]*mana[i]), prefix[j]+int64(skill[j]*mana[i]))
		}

		// Because we have the time earliest that can complete a potion when pass sequentially through all wizard
		// So we must subtract the brewing time to compute the must-be done time of previous wizard
		// reverse iteration to specify the possible earliest time that each wizard must complete their process
		prefix[n-1] = current
		for h := n - 2; h >= 0; h-- {
			prefix[h] = prefix[h+1] - int64(skill[h+1]*mana[i])
		}
		fmt.Println("start: ", prefix[0])
	}

	return prefix[n-1]
}
