package snake

type SController interface {
	//calculates the next Step of the Snake
	NextStep()

	//returns the snake
	GetSnake() Snake
}

type PlayGroundController interface {
	// Starts the Snake
	Start()

	// calculates the next period
	nextPeriod(snake Snake) bool
}
