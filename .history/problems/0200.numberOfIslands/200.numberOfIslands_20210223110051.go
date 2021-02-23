package problems

func checkAround(x, y int, grid [][]byte, visited [][]bool) {

	visited[x][y] = true

	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, d := range direction {
		newX, newY := x+d[0], y+d[1]
		if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == '1' && !visited[newX][newY] {
			checkAround(newX, newY, grid, visited)
		}
	}
}

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				ans++
				checkAround(i, j, grid, visited)
			}
		}
	}
	return ans

}
