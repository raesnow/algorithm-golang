package array

import "fmt"

func carPooling(trips [][]int, capacity int) bool {
	passengers := make([]int, 1001)
	diffs := make([]int, len(passengers)+1)

	for _, trip := range trips {
		diffs[trip[1]] += trip[0]
		diffs[trip[2]] -= trip[0]
	}
	fmt.Printf("diffs: %v", diffs[0:10])

	for i := range passengers {
		if i == 0 {
			passengers[i] = diffs[i]
		} else {
			passengers[i] = passengers[i-1] + diffs[i]
		}
		if passengers[i] > capacity {
			fmt.Printf("passengers: %v, i: %d\n", passengers[0:i+1], i)
			return false
		}
	}
	return true
}
