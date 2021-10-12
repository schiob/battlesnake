package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"log"
	"math/rand"
)

// This function is called when you register your Battlesnake on play.battlesnake.com
// See https://docs.battlesnake.com/guides/getting-started#step-4-register-your-battlesnake
// It controls your Battlesnake appearance and author permissions.
// For customization options, see https://docs.battlesnake.com/references/personalization
// TIP: If you open your Battlesnake URL in browser you should see this data.
func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "schiob",
		Color:      "#3875c9",
		Head:       "default",
		Tail:       "default",
	}
}

// This function is called everytime your Battlesnake is entered into a game.
// The provided GameState contains information about the game that's about to be played.
// It's purely for informational purposes, you don't have to make any decisions here.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}
	myHead := state.You.Body[0] // Coordinates of your head

	boardWidth := state.Board.Width
	if myHead.X == boardWidth-1 {
		possibleMoves["right"] = false
	} else if myHead.X == 0 {
		possibleMoves["left"] = false
	}

	boardHeight := state.Board.Height
	if myHead.Y == boardHeight-1 {
		possibleMoves["up"] = false
	} else if myHead.Y == 0 {
		possibleMoves["down"] = false
	}

	for _, snake := range state.Board.Snakes {
		for _, part := range snake.Body {
			if isLeft(part, myHead) {
				possibleMoves["left"] = false
			} else if isRight(part, myHead) {
				possibleMoves["right"] = false
			} else if isDown(part, myHead) {
				possibleMoves["down"] = false
			} else if isUp(part, myHead) {
				possibleMoves["up"] = false
			}
		}
	}

	// Choose a move from the available safe moves.
	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}

	if len(state.Board.Food) != 0 {
		food := state.Board.Food[0]
		if myHead.X < food.X && possibleMoves["right"] {
			nextMove = "right"
		} else if myHead.X > food.X && possibleMoves["left"] {
			nextMove = "left"
		} else if myHead.Y < food.Y && possibleMoves["up"] {
			nextMove = "up"
		} else if myHead.Y > food.Y && possibleMoves["down"] {
			nextMove = "down"
		}
	}

	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}

func isLeft(target Coord, myHead Coord) bool {
	return target.X == myHead.X-1 && target.Y == myHead.Y
}

func isRight(target Coord, myHead Coord) bool {
	return target.X == myHead.X+1 && target.Y == myHead.Y
}

func isDown(target Coord, myHead Coord) bool {
	return target.Y == myHead.Y-1 && target.X == myHead.X
}

func isUp(target Coord, myHead Coord) bool {
	return target.Y == myHead.Y+1 && target.X == myHead.X
}
