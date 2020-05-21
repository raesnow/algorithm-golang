package main

/*
You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.

The lock initially starts at '0000', a string representing the state of the 4 wheels.

You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.

Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.

Example 1:

Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
Output: 6
Explanation:
A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
because the wheels of the lock become stuck after the display becomes the dead end "0102".

Example 2:

Input: deadends = ["8888"], target = "0009"
Output: 1
Explanation:
We can turn the last wheel in reverse to move from "0000" -> "0009".

Example 3:

Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
Output: -1
Explanation:
We can't reach the target without getting stuck.

Example 4:

Input: deadends = ["0000"], target = "8888"
Output: -1
Note:

The length of deadends will be in the range [1, 500].
target will not be in the list deadends.
Every string in deadends and the string target will be a string of 4 digits from the 10,000 possibilities '0000' to '9999'.
*/

func openLock(deadends []string, target string) int {
	queue := make([]string, 0)
	visited := make(map[string]bool)
	for _, v := range deadends {
		visited[v] = true
	}
	stepCount := 0

	queue = append(queue, "0000")
	for len(queue) != 0 {
		size := len(queue)
		newQueue := make([]string, 0)

		for i := 0; i < size; i++ {
			value := queue[i]
			if value == target {
				return stepCount
			}
			if _, ok := visited[value]; ok {
				continue
			} else {
				visited[value] = true
			}

			result := turnOne(value)
			newQueue = append(newQueue, result...)
		}
		queue = newQueue
		stepCount++
	}
	return -1
}

func turnOne(value string) (result []string) {
	for i := range value {
		newValue := make([]byte, len(value))
		copy(newValue, []byte(value))

		up := value[i] - 1
		if value[i] == '0' {
			up = '9'
		}
		newValue[i] = up
		result = append(result, string(newValue))

		down := value[i] + 1
		if value[i] == '9' {
			down = '0'
		}
		newValue[i] = down
		result = append(result, string(newValue))
	}
	return result
}

func openLock1(deadends []string, target string) int {
	queue1 := make(map[string]bool)
	queue2 := make(map[string]bool)
	visited := make(map[string]bool)
	for _, v := range deadends {
		visited[v] = true
	}
	stepCount := 0

	queue1["0000"] = true
	queue2[target] = true
	for len(queue1) != 0 && len(queue2) != 0 {
		newQueue := make(map[string]bool)

		for value := range queue1 {
			if _, ok := queue2[value]; ok {
				return stepCount
			}
			if _, ok := visited[value]; ok {
				continue
			} else {
				visited[value] = true
			}

			result := turnOne(value)
			for _, v := range result {
				newQueue[v] = true
			}
		}
		queue1 = newQueue
		queue1, queue2 = queue2, queue1
		stepCount++
	}
	return -1
}
