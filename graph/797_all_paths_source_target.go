package graph

var result = make([][]int, 0)

func allPathsSourceTarget(graph [][]int) [][]int {
	if len(graph) == 0 {
		return nil
	}

	traverse(graph, 0, []int{})
	return result
}

func traverse(graph [][]int, node int, path []int) {
	path = append(path, node)

	if node == len(graph)-1 {
		result = append(result, path)
	}

	for _, neighbour := range graph[node] {
		traverse(graph, neighbour, path)
	}
	path = path[:len(path)-1]
}
