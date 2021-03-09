package problems

func rotate(matrix [][]int) {
	n := len(matrix)
	l := n - 1
	var tmp int
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2+n%2; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[l-j][i]
			matrix[l-j][i] = matrix[l-i][l-j]
			matrix[l-i][l-j] = matrix[j][l-i]
			matrix[j][l-i] = tmp
		}
	}
}
