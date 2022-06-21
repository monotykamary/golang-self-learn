package main

import "fmt"

// For matrices, if you see a right turn, check left

func generatePascalsTriangle(numRows int) [][]int {
	/*
		Create a int matrix with the length (height) of the matrix equal to numRows.
	*/
	var outputs = make([][]int, numRows)

	/*
		Create a int matrix with the length (height) of the matrix equal to numRows
	*/
	for i := 0; i < numRows; i++ {
		/*
			Create an int slice for each row with the length (width) of the row equal to the index + 1
			(as we start i with 0)
		*/
		outputs[i] = make([]int, i+1)

		/*
			All edges of a row start at 1; we'll optimize it per row here
		*/
		outputs[i][0] = 1
		outputs[i][i] = 1

		/*
			Loop through each item inside the row i. We start at 1 to avoid j-1 being out of index.
		*/
		for j := 1; j < i; j++ {
			/*
				j-1 is our top left adjacent number.
				j is our top right adjacent number.
			*/
			outputs[i][j] = outputs[i-1][j-1] + outputs[i-1][j]
		}
	}

	return outputs
}

func main() {
	triangle := generatePascalsTriangle(5)
	fmt.Printf("%v", triangle)
}
