package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
)

type Board [3][3]string

var cursorX, cursorY int

func initializeBoard() Board {
	var board Board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "_"
		}
	}
	cursorX, cursorY = 0, 0
	return board
}

func printBoard(board Board) {
	clearScreen()
	fmt.Print("Use arrow keys to move the cursor. Press Space to place your mark. \n\n")
	fmt.Println("   0   1   2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			if i == cursorX && j == cursorY {
				fmt.Printf("[%s] ", board[i][j])
			} else {
				fmt.Printf(" %s  ", board[i][j])
			}
		}
		fmt.Print("\n\n")

	}
}

func clearScreen() {
	cmd := exec.Command("clear") // for Linux/Unix
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func checkWin(player string, board Board) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func isBoardFull(board Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "_" {
				return false
			}
		}
	}
	return true
}

func switchPlayer(currentPlayer *int) {
	if *currentPlayer == 1 {
		*currentPlayer = 2
	} else {
		*currentPlayer = 1
	}
}

func main() {
	var currentPlayer int
	currentPlayer = 1
	var board Board
	board = initializeBoard()
	keyboard.Open()
	defer keyboard.Close()

	printBoard(board)

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			return
		}
		if key == keyboard.KeyArrowUp {
			if cursorX > 0 {
				cursorX--
			}
		} else if key == keyboard.KeyArrowDown {
			if cursorX < 2 {
				cursorX++
			}
		} else if key == keyboard.KeyArrowLeft {
			if cursorY > 0 {
				cursorY--
			}
		} else if key == keyboard.KeyArrowRight {
			if cursorY < 2 {
				cursorY++
			}
		} else if key == keyboard.KeySpace {
			if board[cursorX][cursorY] == "_" {
				if currentPlayer == 1 {
					board[cursorX][cursorY] = "X"
				} else {
					board[cursorX][cursorY] = "O"
				}

				if checkWin(board[cursorX][cursorY], board) {
					printBoard(board)
					fmt.Printf("Player %d wins!\n", currentPlayer)
					return
				} else if isBoardFull(board) {
					printBoard(board)
					fmt.Println("It's a draw!")
					return
				}

				switchPlayer(&currentPlayer)
			}
		} else if char == 'q' || char == 'Q' { // Condition to quit the game (Using keyboard strokes hence, CMD + C won't work.)
			return
		}
		printBoard(board)
	}
}
