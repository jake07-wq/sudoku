package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 9 {
		fmt.Println("Error")
		return
	}
	var board [9][9]rune
	for i, rowStr := range args {
		if len(rowStr) != 9 {
			fmt.Println("Error")
			return
		}
		for j, char := range rowStr {
			board[i][j] = char
		}
	}
	if solve(&board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}

func solve(board *[9][9]rune) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				for num := '1'; num <= '9'; num++ {
					if isValid(board, r, c, num) {
						board[r][c] = num
						if solve(board) {
							return true
						}
						board[r][c] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board *[9][9]rune, r, c int, num rune) bool {
	for i := 0; i < 9; i++ {
		if board[r][i] == num || board[i][c] == num ||
			board[3*(r/3)+i/3][3*(c/3)+i%3] == num {
			return false
		}
	}
	return true
}

func printBoard(board [9][9]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", board[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
