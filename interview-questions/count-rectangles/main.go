package main

import "fmt"

// For matrices, if you see a right turn, check left

func countRectangles(rectangles [][]int) int {
	count := 0

	// Range over rectangles to get its rows and index
	for i, row := range rectangles {

		// Range over values to get its values and index
		for j, value := range row {
			// Determine if the adjacent numbers are 0; if they are, add to count; else we'll continue traversing
			// Check if the current number is 0
			if value == 0 {
				continue
			}

			// Check if we're not at the start of the row and check if the number on top of us is 1
			if i > 0 {
				upValue := rectangles[i-1][j]
				if upValue == 1 {
					continue
				}
			}

			// Check if we're not at the start of the row and check if the number on left of us is 1
			if j > 0 {
				leftValue := rectangles[i][j-1]
				if leftValue == 1 {
					continue
				}
			}

			/*
			   When all checks previous checks for 1 don't exist, we've reached the edge of the rectangle (that is 0),
			   meaning that what we have is probably a rectangle; increase our count here
			*/
			count++
		}
	}
	return count
}

func main() {
	arr := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	count := countRectangles(arr)
	fmt.Printf("%v", count)
}
