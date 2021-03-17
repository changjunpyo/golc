package problems

func searchMatrix(matrix [][]int, target int) bool {
	i := len(matrix) - 1
	j := 0
	m := len(matrix[0])
	for i >= 0 && j < m {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false

}
