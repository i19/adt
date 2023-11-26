package leetcode

import "math"

// 258. 逃离火灾
// https://leetcode.cn/problems/escape-the-spreading-fire/description/
func maximumMinutes(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	fireTime := 0

	for grid[rows-1][cols-1] != 1 {
		fireTime++
		newFireCount := 0
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				if grid[row][col] == 1 {
					newFireCount += setFire(grid, row, col)
				}
			}
		}

		if newFireCount == 0 {
			break
		}
	}

	if grid[rows-1][cols-1] == 1 {
		return int(math.Pow(10, 9))
	}

	t := fireTime - (rows + cols - 1)
	println("hello world")
	if t > 0 {
		return t
	} else {
		return -1
	}
}

func setFire(grid [][]int, pRow, pCol int) int {
	rows, cols := len(grid), len(grid[0])
	upperRow, lowerRow := pRow-1, pRow+1
	leftCol, rightCol := pCol-1, pCol+1
	upperRow = max(0, upperRow)
	lowerRow = min(rows-1, lowerRow)
	leftCol = max(0, leftCol)
	rightCol = min(cols-1, rightCol)

	count := 0
	for row := lowerRow; row <= upperRow; row++ {
		for col := leftCol; col <= rightCol; col++ {
			if grid[row][col] == 0 {
				grid[row][col] = 1
				count++
			}
		}
	}
	return count
}
