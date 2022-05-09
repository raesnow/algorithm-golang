package main

var permuteResult [][]int

func permute1(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	permuteResult = make([][]int, 0)
	permuteBacktrack([]int{}, nums, len(nums))
	return permuteResult
}

func permuteBacktrack(path []int, choices []int, length int) {
	if len(path) == length {
		result := make([]int, len(path))
		copy(result, path)
		permuteResult = append(permuteResult, result)
		return
	}

	for i, choice := range choices {
		path = append(path, choice)
		newChoices := make([]int, 0)
		newChoices = append(newChoices, choices[0:i]...)
		newChoices = append(newChoices, choices[i+1:]...)
		permuteBacktrack(path, newChoices, length)
		path = path[:len(path)-1]
	}
}
