package main

import "fmt"

func initFlodfill(state GameState) [][]int {
	arr := make([][]int, state.Board.Height)
	for i := range arr {
		arr[i] = make([]int, state.Board.Width)
	}

	// Add snakes
	for _, snake := range state.Board.Snakes {
		for _, part := range snake.Body {
			arr[part.Y][part.X] = -1
		}
	}

	colorCont := 1
	for y := 0; y < len(arr); y++ {
		for x := 0; x < len(arr); x++ {
			if arr[y][x] == 0 {
				flodFill(arr, x, y, colorCont)
				colorCont++
			}
		}
	}

	printArray(arr)
	return arr
}

func flodFill(arr [][]int, x, y int, color int) {
	// Base
	if x < 0 || x >= len(arr) {
		return
	}
	if y < 0 || y >= len(arr) {
		return
	}

	// Snake
	if arr[y][x] != 0 {
		return
	}

	// Color
	arr[y][x] = color
	flodFill(arr, x, y+1, color)
	flodFill(arr, x, y-1, color)
	flodFill(arr, x-1, y, color)
	flodFill(arr, x+1, y, color)
}

func printArray(arr [][]int) {
	for i := range arr {
		for _, val := range arr[len(arr)-1-i] {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}
