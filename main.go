package main

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-asiansneverdie/snake"
)

func main() {
	playground := snake.NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.CreateOuterBorders()
	s := snake.NewSnake(4)
	playground.CreateSnake(s)
	snakeController := snake.NewSnakeController(playground, s)
	//Speed in MillieSeconds
	gameController := snake.NewPGController(playground, snakeController, 50)

	gameController.Start()

}
