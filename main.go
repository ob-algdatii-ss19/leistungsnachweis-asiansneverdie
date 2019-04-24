package main

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-asiansneverdie/snake"
)

func main() {
	playground := snake.NewPlayground()
	playground.CreateEmptyPlayground(40, 30)
	playground.CreateOuterBorders()
	s := snake.NewSnake(4)
	snakeController := snake.NewSnakeController(&playground, s)
	//Speed in MillieSeconds
	gameController := snake.NewPGController(playground, snakeController, 1000)

	gameController.Start()

}
