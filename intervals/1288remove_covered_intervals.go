package main

import "sort"

/*
Given a list of intervals, remove all intervals that are covered by another interval in the list.

Interval [a,b) is covered by interval [c,d) if and only if c <= a and b <= d.

After doing so, return the number of remaining intervals.

Example 1:

Input: intervals = [[1,4],[3,6],[2,8]]
Output: 2
Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
Example 2:

Input: intervals = [[1,4],[2,3]]
Output: 1
Example 3:

Input: intervals = [[0,10],[5,12]]
Output: 2
Example 4:

Input: intervals = [[3,10],[4,10],[5,11]]
Output: 2
Example 5:

Input: intervals = [[1,2],[1,4],[3,4]]
Output: 1


Constraints:

1 <= intervals.length <= 1000
intervals[i].length == 2
0 <= intervals[i][0] < intervals[i][1] <= 10^5
All the intervals are unique.
*/

func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Sort(intervalType(intervals))

	end := intervals[0][1]
	covered := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] <= end {
			covered++
		} else {
			end = intervals[i][1]
		}
	}
	return len(intervals) - covered
}

type intervalType [][]int

func (v intervalType) Len() int {
	return len(v)
}

func (v intervalType) Less(i, j int) bool {
	if v[i][0] < v[j][0] {
		return true
	} else if v[i][0] == v[j][0] && v[i][1] > v[j][1] {
		return true
	}
	return false
}

func (v intervalType) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
