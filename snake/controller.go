package snake

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
