package main

func eggDrop(eggsCount, floorCount int) int {
	cache := make([][]int, floorCount+1)
	for i := range cache {
		cache[i] = make([]int, eggsCount+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	return dp(cache, eggsCount, floorCount)
}

func dp(cache [][]int, eggsCount, floorCount int) int {
	if eggsCount == 1 {
		return floorCount
	}
	if floorCount == 0 {
		return 0
	}

	if cache[floorCount][eggsCount] != -1 {
		return cache[floorCount][eggsCount]
	}

	result := 0
	for i := 1; i <= floorCount; i++ {
		broken := dp(cache, eggsCount-1, i-1)
		notBroken := dp(cache, eggsCount, floorCount-i)
		value := broken
		if notBroken > broken {
			value = notBroken
		}
		if result == 0 || result > value+1 {
			result = value + 1
		}
	}
	cache[floorCount][eggsCount] = result
	return result
}

func eggDrop1(eggsCount, floorCount int) int {
	cache := make([][]int, floorCount+1)
	for i := range cache {
		cache[i] = make([]int, eggsCount+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	return dp(cache, eggsCount, floorCount)
}

func dp1(cache [][]int, eggsCount, floorCount int) int {
	if eggsCount == 1 {
		return floorCount
	}
	if floorCount == 0 {
		return 0
	}

	if cache[floorCount][eggsCount] != -1 {
		return cache[floorCount][eggsCount]
	}

	result := 0
	left, right := 1, floorCount
	for left <= right {
		mid := left + (right-left)/2
		broken := dp(cache, eggsCount-1, mid-1)
		notBroken := dp(cache, eggsCount, floorCount-mid)
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

	cache[floorCount][eggsCount] = result
	return result
}

func eggDrop2(eggsCount, floorCount int) int {
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

func eggDrop3(eggsCount, floorCount int) int {
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
