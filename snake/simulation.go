package snake

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/mohae/deepcopy"
)
// TODO: correct input
func Simulate(sc *SimpleSnakeController, move []DIRECTION, snakeLength int) bool {
	//nolint
	if len(move) < 1 {
		sc.Pg.Print()
		return false
	}
	sc.moveSnakeToFood(move)
	snake := sc.Snake

	if sc.Pg.GetContent(snake.Head.X, snake.Head.Y) == BORDER {
		sc.Pg.Print()
		return false
	} else if sc.Pg.GetContent(snake.Head.X, snake.Head.Y) == FOOD {
		sc.Pg.setRandomFood()
	}
	sc.Pg.DeleteSnake()
	sc.Pg.CreateSnake(snake)

	if snakeLength < 1 {
		//sc.Pg.Print()
		return true
	} else {
		return Simulate(sc, GetDirections(sc.Pg.GetPlayGround()), snakeLength - 1)
	}
	return Simulate(sc, GetDirections(sc.Pg.GetPlayGround()), depth, snakeLength-1)
}

func (sc *SimpleSnakeController) GetNextMovableFoodDirection(move []DIRECTION) DIRECTION {
	dir := move[0]
	var x, y = sc.Pg.GetFood()
	//nolint
	if sc.Snake.Head.X < x && contains(move, RIGHT) {
		return RIGHT
	} else if sc.Snake.Head.X > x && contains(move, LEFT) {
		return LEFT
	} else if sc.Snake.Head.Y < y && contains(move, DOWN) {
		return DOWN
	} else if contains(move, UP) {
		return UP
	}
	return dir
}

func Remove(move []DIRECTION, element DIRECTION) []DIRECTION {
	movecopy := []DIRECTION{}
	for ele := range move {
		if DIRECTION(ele) != element {
			movecopy = append(movecopy, DIRECTION(ele))
		}
	}
	return movecopy
}
