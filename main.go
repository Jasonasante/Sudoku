package main

import (
	"os"

	"github.com/01-edu/z01"
)

var sudo [9][9]rune //This is saying that everything inside the 9 by 9 table is a rune

func isBoardValid(guess rune, row, col int) bool { //Create variable type rune, , row and
	//column are both integers and bool means true or false.
	//checking if the number in the column is valid
	for j := 0; j < 9; j++ { // this is for the column/
		if sudo[row][j] == guess { //if at number
			return false
		}
	}
	//Checking if the number in the row is valid
	for i := 0; i < 9; i++ {
		if sudo[i][col] == guess {
			return false
		}
	}
	//Checking if the box is valid
	//This is the way to get the initial position to iterate over the box
	row_start := (row / 3) * 3
	col_start := (col / 3) * 3
	for i := row_start; i < row_start+3; i++ {
		for j := col_start; j < col_start+3; j++ {
			if sudo[i][j] == guess {
				return false
			}
		}
	}
	return true
}

//Is going look for next dot, that means there is no number there
func findNextEmpty() (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudo[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func solveSudoku() bool {
	row, col := findNextEmpty()
	if row == -1 {
		return true
	}
	//Is going to check if the number is valid, and is going to assign it on the right position
	for i := 1; i < 10; i++ {
		if isBoardValid(rune(i+48), row, col) == true {
			sudo[row][col] = rune(i + 48)
			if solveSudoku() == true {
				return true
			}
		}
		//this is going to allow to backtrack
		sudo[row][col] = '.'
	}
	return false
}

func messageError() {
	error := "Error"
	for i := 0; i < len(error); i++ {
		z01.PrintRune(rune(error[i]))
	}
	z01.PrintRune('\n')
}

func main() {
	//Holds the arguments to the program.
	args := os.Args[1:]
	//If the arguments are not 9, error
	if len(args) != 9 {
		messageError()
		return
	}
	//This loop traverses vertically
	for i := 0; i < 9; i++ {
		if len(args[i]) != 9 {
			messageError()
			return
		}
		//This loop traverses horizontally
		for j := 0; j < 9; j++ {
			//If a number
			if args[i][j] > '0' && args[i][j] <= '9' {
				//if the number is valid asccess it
				if isBoardValid(rune(args[i][j]), i, j) == true {
					sudo[i][j] = rune(args[i][j])
				} else {
					messageError()
					return
				}
				//if not a number
			} else {
				//if not a dot, error
				if args[i][j] != '.' {
					messageError()
					return
				}
				//if a dot, assign the dot to the sudoku
				sudo[i][j] = '.'
			}
		}
	}
	//print sudoku
	if solveSudoku() == true {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				z01.PrintRune(sudo[i][j])
				if j != 8 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
		//print error
	} else {
		messageError()
		return
	}
}
