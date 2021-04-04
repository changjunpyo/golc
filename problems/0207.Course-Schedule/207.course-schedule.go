package problems

func canFinish(numCourses int, prerequisites [][]int) bool {
	edge := make([][]int, numCourses)
	visited := make([]bool, numCourses)
	inComming := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		edge[i] = make([]int, 0)
	}
	for i := 0; i < len(prerequisites); i++ {
		edge[prerequisites[i][0]] = append(edge[prerequisites[i][0]], prerequisites[i][1])
		inComming[prerequisites[i][1]] += 1
	}
	for i := 0; i < numCourses; i++ {
		if inComming[i] == 0 {
			DFS(edge, visited, inComming, i)
		}
	}
	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}

func DFS(edge [][]int, visited []bool, incomming []int, node int) {
	if visited[node] {
		return
	}
	visited[node] = true
	for _, next := range edge[node] {
		incomming[next]--
		if incomming[next] == 0 {
			DFS(edge, visited, incomming, next)
		}
	}
}
