// https://leetcode.com/problems/01-matrix/
package main

import "fmt"

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(updateMatrix(mat))
}

// solution 1: BFS (breadth-first search)
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	q := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != 0 {
				mat[i][j] = -1
			} else {
				q = append(q, []int{i, j})
			}
		}
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, dir := range dirs {
			x, y := cur[0]+dir[0], cur[1]+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n || mat[x][y] != -1 {
				continue
			}
			mat[x][y] = mat[cur[0]][cur[1]] + 1
			q = append(q, []int{x, y})
		}
	}
	return mat
}

// solution 2: DP (dynamic programming)
func updateMatrix2(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = m + n
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
				}
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
				}
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] != 0 {
				if i < m-1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
				if j < n-1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
			}
		}
	}
	return dp
}
