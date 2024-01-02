package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
)

var board [3][3]string
var currentPlayer int
var cursorX, cursorY int

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "_"
		}
	}
	currentPlayer = 1
	cursorX, cursorY = 0, 0
}

func printBoard() {
	clearScreen()
	fmt.Println("Use arrow keys to move the cursor. Press Space to place your mark. \n")
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
		fmt.Println("\n")

	}
}

func clearScreen() {
	cmd := exec.Command("clear") // for Linux/Unix
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func checkWin(player string) bool {
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

func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "_" {
				return false
			}
		}
	}
	return true
}

func switchPlayer() {
	if currentPlayer == 1 {
		currentPlayer = 2
	} else {
		currentPlayer = 1
	}
}

func main() {
	initializeBoard()
	keyboard.Open()
	defer keyboard.Close()

	printBoard()

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

				if checkWin(board[cursorX][cursorY]) {
					printBoard()
					fmt.Printf("Player %d wins!\n", currentPlayer)
					return
				} else if isBoardFull() {
					printBoard()
					fmt.Println("It's a draw!")
					return
				}

				switchPlayer()
			}
		} else if char == 'q' || char == 'Q' { // Condition to quit the game (Using keyboard strokes hence, CMD + C won't work.)
			return 
		}
		printBoard()
	}
}
