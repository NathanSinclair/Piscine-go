package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 10 { // if the length of the argumatents including os.Args[0] this means that we have valid input else return error
		args := make([][]string, 9) // Here we are creating a 2d splice of length 9. This 1 column and 9 rows
		// This looks like: [[] [] [] [] [] [] [] [] []]

		/* Here we are taking each argument from os.Args starting from os.Args[1] and adding it to
		   to the args array
		   Right now it looks like this:
		   [.96.4...1]
		   [1...6...4]
		   [5.481.39.]
		   [..795..43]
		   [.3..8....]
		   [4.5.23.18]
		   [.1.63..59]
		   [.59.7.83.]
		   [..359...7]
		   (We start from os.Args[1] because os.Args[0] is the programme name)
		*/
		for i := 1; i < len(os.Args); i++ {
			args[i-1] = append(args[i-1], os.Args[i])
		}
		// iterate through args and convert each row to a splice of runes and add that to 2d runes
		// recieved help from : https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go
		board := make([][]rune, 9) // Here we are creating a 2d array called board the default is 9 rows and 1 column
		for i := 0; i < 9; i++ {
			board[i] = make([]rune, 9) // Now we are iterating through each row and adding 9 cols
		}
		// fmt.Println(board)
		for i := 0; i < len(args); i++ {
			runeSplice := []rune(args[i][0]) // convert each string to a runesplice.
			for j, r := range runeSplice {
				board[i][j] = r
			}

		}
		if solvingPuzzle(board) { // if we have a solution call the function that prints the solution
			printingSolution(board)
		} else { // this means that there is no valid solution
			z01.PrintRune('E')
			z01.PrintRune('r')
			z01.PrintRune('r')
			z01.PrintRune('o')
			z01.PrintRune('r')
			z01.PrintRune('\n')
		}
	} else {
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('\n')
	}

}

// function to print the final board with the solution.
// This function is only called when we have a valid solution so if solving puzzle returns true
func printingSolution(board [][]rune) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			z01.PrintRune(board[i][j])
			if j < len(board[i])-1 {
				z01.PrintRune(' ')
			}

		}
		z01.PrintRune('\n')
	}
}

/* this function checks if the number we want to place in a given row is already there. Returns true if it is already there and false if it isn't
puzzle  = We pass in the board we are trying to check
num = This is the number we want to check
row = the index of the row we want to check

*/
func checkNumberInRow(puzzle [][]rune, num rune, row int) bool {
	for i := 0; i < 9; i++ {
		if puzzle[row][i] == num { // here we keep the row number constant and iterate through the columns (i.e. moving left to right)
			return true
		}
	}
	return false
}

// this function checks if the number we want to place in a given column is already there
// puzzle  = We pass in the board we are trying to check
// col = the index of the col we want to check

func checkNumberInCol(puzzle [][]rune, num rune, col int) bool {
	for i := 0; i < 9; i++ {
		if puzzle[i][col] == num { // here we are keeping the column number constant and iterate through the rows (i.e moving up and down)
			return true
		}
	}
	return false
}

// Given any index we are in we calculate the co-ordinates of the top left corner and of the top right corner of the 3x3 subsection
// and iterate through this section to see if the number repeats itself in this box
func checkNumberIn3x3(puzzle [][]rune, num rune, row int, col int) bool {
	var topLeftCornerRow int = row - (row % 3)
	var topLeftCornerCol int = col - (col % 3)
	for i := topLeftCornerRow; i < topLeftCornerRow+3; i++ {
		for j := topLeftCornerCol; j < topLeftCornerCol+3; j++ {
			if puzzle[i][j] == num {
				return true
			}
		}
	}
	return false
}

// if the number is NOT in the row AND NOT in the column AND NOT in the 3x3 matrix return true
func isValid(puzzle [][]rune, num rune, row int, col int) bool {
	return !checkNumberInRow(puzzle, num, row) && !checkNumberInCol(puzzle, num, col) && !checkNumberIn3x3(puzzle, num, row, col)
}
func solvingPuzzle(puzzle [][]rune) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if puzzle[row][col] == '.' {
				// if we encounter a dot, let's try the numbers from 1 to 9 to see if we can get find a valid number that works
				for try := '1'; try <= '9'; try++ {
					if isValid(puzzle, try, row, col) { // if this is true place the number and replace the dot
						puzzle[row][col] = try
						if solvingPuzzle(puzzle) {
							// now we recursively calel our solvingPuzzle() function looking for the next '.'
							return true // this means we are done
						} else {
							puzzle[row][col] = '.'
							// even though the number is valid from the previous if statement
							// we can't solve the rest of the board so we need to reset it back to '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}
