package main

import (
	"fmt"
)

// For matrices, if you see a right turn, check left

func isValidSudoku(board [][]byte) bool {
	/*
		Create a cache of matrices that return a bool. This is so whenever we hit a value
		that is in our cache and it returns true, we've hit a duplicate number. We will use
		this as our ghetto lookup table.

		The size of the matrices are 9x9 (first 9 representing the row/column, second 9
		holding the numbers we cache from 1 - 9).
	*/
	row := [9][9]bool{}
	col := [9][9]bool{}

	/*
		We create box here differently as we are checking 3x3; our matrix here will be
		3x3x9 (first 3 is our rows, second 3 is our columns, succeeding 9 is holding the
		numbers we cache from 1 - 9)
	*/
	box := [3][3][9]bool{}

	/*
		Instantiate the arrays/matrices inside our arrays.
	*/
	for i := 0; i < 9; i++ {
		row[i] = [9]bool{}
		col[i] = [9]bool{}

		if i < 3 {
			box[i] = [3][9]bool{}

			for j := 0; j < 3; j++ {
				box[i][j] = [9]bool{}
			}
		}
	}

	/*
		Prepare 2 loops for rows and columns to begin traversing through our Sudoku board.
	*/
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			/*
				Check if the board contains values between 1 and 9.
			*/
			if board[r][c] >= '1' && board[r][c] <= '9' {
				/*
					Find what the current digit in bytes. Once we have that, subtract the value by
					`0` (0x30) and convert the value into an integer so we can use it to access our
					arrays.
				*/
				digit := board[r][c]
				ref := int(digit-'0') - 1

				/*
					To get flatten the row and column values into an independent 3x3 box, we need to
					know that, by index, our grid will be sorted as such:

					[0]: {0, 1, 2}
					[1]: {3, 4, 5}
					[2]: {6, 7, 8}

					With the nuance Golang has with integers, we can take advantage of it to ignore
					remainders and give us a value equal to the left-most value in the array / 3
				*/
				rBox := r / 3
				cBox := c / 3

				/*
					If any of the values exist between the row, col, and box caches, we've hit a duplicate.
					We can immediately return false here.
				*/
				if row[r][ref] {
					return false
				}

				if col[c][ref] {
					return false
				}

				if box[rBox][cBox][ref] {
					return false
				}

				/*
					Else, map the number reference to our cache and assign the value to true.
				*/
				row[r][ref] = true
				col[c][ref] = true
				box[rBox][cBox][ref] = true
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	isValid := isValidSudoku(board)
	fmt.Printf("%v", isValid)
}
