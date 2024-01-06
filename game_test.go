package main

import (
	"fmt"
	"testing"
)

func TestCheckWin(t *testing.T) {
	board := [3][3]string{
		{"X", "X", "X"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	if !checkWin("X", board) {
		t.Error("Horizontal win error (Player X)")

	} else {
		t.Log("Horizontal win success (Player X)")
	}

	board = [3][3]string{
		{"O", "_", "_"},
		{"O", "_", "_"},
		{"O", "_", "_"},
	}

	if !checkWin("O", board) {
		t.Error("Vertical win error (Player O)")
	} else {
		t.Log("Vertical win success (Player O)")
	}

	board = [3][3]string{
		{"_", "_", "X"},
		{"_", "X", "_"},
		{"X", "_", "_"},
	}

	if !checkWin("X", board) {
		t.Error("Diagonal win error (Player X)")
	} else {
		t.Log("Diagonal win success (Player X)")
	}

	board = [3][3]string{
		{"O", "X", "O"},
		{"X", "O", "X"},
		{"X", "O", "X"},
	}

	if checkWin("X", board) || checkWin("O", board) {
		t.Error("Win condition error (Player X and O)")
	} else {
		t.Log("Win condition success (Player X and O)")
	}

	defer func() {
		if t.Failed() {
			return
		}
		fmt.Println("TestCheckWin -- PASS")
	}()
}

func TestIsBoardFull(t *testing.T) {
	board := Board{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	if isBoardFull(board) {
		t.Error("Empty board detection error")
	} else {
		t.Log("Empty board detection success")
	}

	board = Board{
		{"X", "_", "O"},
		{"O", "X", "X"},
		{"X", "O", "_"},
	}

	if isBoardFull(board) {
		t.Error("Partially filled board detection error")
	} else {
		t.Log("Partially filled board detection success")
	}

	board = Board{
		{"X", "X", "O"},
		{"O", "O", "X"},
		{"X", "X", "O"},
	}

	if !isBoardFull(board) {
		t.Error("Fully filled board detection error")
	} else {
		t.Log("Fully filled board detection success")
	}

	defer func() {
		if t.Failed() {
			return
		}
		fmt.Println("TestIsBoardFull -- PASS")
	}()
}
