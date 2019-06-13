package snake

import Snake "github.com/ob-algdatii-ss19/leistungsnachweis-asiansneverdie/snake/snake"

type SController interface {
	//calculates the next Step of the Snake
	NextStep()

	//returns the snake
	GetSnake() Snake
	addTail() Snake
}

type PlayGroundController interface {
	// Starts the Snake
	Start()
}
