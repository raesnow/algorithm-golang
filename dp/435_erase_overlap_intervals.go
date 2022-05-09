package main

import "sort"

type Intervals [][]int

func (ints Intervals) Len() int {
	return len(ints)
}

func (ints Intervals) Less(i, j int) bool {
	return ints[i][1] < ints[j][1]
}

func (ints Intervals) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Sort(Intervals(intervals))

	preEnd := intervals[0][1]
	noOverlapCount := 1
	for _, interval := range intervals[1:] {
		if interval[0] >= preEnd {
			noOverlapCount++
			preEnd = interval[1]
		}
	}
	return len(intervals) - noOverlapCount
}
