package main

/*
You are given K eggs, and you have access to a building with N floors from 1 to N.

Each egg is identical in function, and if an egg breaks, you cannot drop it again.

You know that there exists a floor F with 0 <= F <= N such that any egg dropped at a floor higher than F will break,
and any egg dropped at or below floor F will not break.

Each move, you may take an egg (if you have an unbroken one) and drop it from any floor X (with 1 <= X <= N).

Your goal is to know with certainty what the value of F is.

What is the minimum number of moves that you need to know with certainty what F is, regardless of the initial value of F?


Example 1:

Input: K = 1, N = 2
Output: 2
Explanation:
Drop the egg from floor 1.  If it breaks, we know with certainty that F = 0.
Otherwise, drop the egg from floor 2.  If it breaks, we know with certainty that F = 1.
If it didn't break, then we know with certainty F = 2.
Hence, we needed 2 moves in the worst case to know what F is with certainty.
Example 2:

Input: K = 2, N = 6
Output: 3
Example 3:

Input: K = 3, N = 14
Output: 4


Note:

1 <= K <= 100
1 <= N <= 10000
*/

var sedCache [][]int

func superEggDrop(K int, N int) int {
	sedCache = make([][]int, N+1)
	for i := range sedCache {
		sedCache[i] = make([]int, K+1)
		for j := range sedCache[i] {
			sedCache[i][j] = -1
		}
	}

	return dp(K, N)
}

func dp(eggsCount, floorCount int) int {
	if eggsCount == 1 {
		return floorCount
	}
	if floorCount == 0 {
		return 0
	}

	if sedCache[floorCount][eggsCount] != -1 {
		return sedCache[floorCount][eggsCount]
	}

	result := 0
	for i := 1; i <= floorCount; i++ {
		broken := dp(eggsCount-1, i-1)
		notBroken := dp(eggsCount, floorCount-i)
		value := broken
		if notBroken > broken {
			value = notBroken
		}
		if result == 0 || result > value+1 {
			result = value + 1
		}
	}
	sedCache[floorCount][eggsCount] = result
	return result
}

func superEggDrop1(K int, N int) int {
	sedCache = make([][]int, N+1)
	for i := range sedCache {
		sedCache[i] = make([]int, K+1)
		for j := range sedCache[i] {
			sedCache[i][j] = -1
		}
	}

	return dp1(K, N)
}

func dp1(eggsCount, floorCount int) int {
	if eggsCount == 1 {
		return floorCount
	}
	if floorCount == 0 {
		return 0
	}

	if sedCache[floorCount][eggsCount] != -1 {
		return sedCache[floorCount][eggsCount]
	}

	result := 0
	left, right := 1, floorCount
	for left <= right {
		mid := left + (right-left)/2
		broken := dp1(eggsCount-1, mid-1)
		notBroken := dp1(eggsCount, floorCount-mid)
		if broken == notBroken {
			result = broken + 1
		} else if broken > notBroken {
			if result == 0 || result > broken+1 {
				result = broken + 1
			}
			right = mid - 1
		} else {
			if result == 0 || result > notBroken+1 {
				result = notBroken + 1
			}
			left = mid + 1
		}
	}

	sedCache[floorCount][eggsCount] = result
	return result
}

func superEggDrop2(eggsCount, floorCount int) int {
	cache := make([][]int, eggsCount+1)
	for i := range cache {
		cache[i] = make([]int, floorCount+1)
	}

	i := 0
	for cache[eggsCount][i] < floorCount {
		i++
		for j := 1; j <= eggsCount; j++ {
			cache[j][i] = cache[j][i-1] + cache[j-1][i-1] + 1
		}
	}
	return i
}

func superEggDrop3(eggsCount, floorCount int) int {
	cache := make([]int, eggsCount+1)

	i := 0
	for cache[eggsCount] < floorCount {
		i++
		newCache := make([]int, eggsCount+1)
		for j := 1; j <= eggsCount; j++ {
			newCache[j] = cache[j] + cache[j-1] + 1
		}
		cache = newCache
	}
	return i
}
