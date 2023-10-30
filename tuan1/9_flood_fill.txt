// https://leetcode.com/problems/flood-fill/
package main

import "fmt"

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr := 1
	sc := 1
	newColor := 2
	floodFill(image, sr, sc, newColor)
	fmt.Println(image)
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	var origColor = image[sr][sc]
	if origColor == color {
		return image
	}
	fill(image, sr, sc, color, origColor)
	return image
}

func fill(image [][]int, sr int, sc int, color int, origColor int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != origColor {
		return
	}
	image[sr][sc] = color
	fill(image, sr-1, sc, color, origColor)
	fill(image, sr+1, sc, color, origColor)
	fill(image, sr, sc-1, color, origColor)
	fill(image, sr, sc+1, color, origColor)
}
