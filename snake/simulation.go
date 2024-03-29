package snake

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/mohae/deepcopy"
)

func Simulate(sc *SimpleSnakeController, move []DIRECTION, depth, snakeLength int) bool {
	//nolint
	if len(move) < 1 {
		return false
	} else if len(move) == 2 && depth < 3 {
		copied := deepcopy.Copy(sc)
		copyController := copied.(*SimpleSnakeController)
		err := copier.Copy(copyController.Pg, sc.Pg)
		if err != nil {
			fmt.Println("Error occurred while copying")
		}
		duplicate := copyController.Pg.CopyPlayGround(copyController.Pg.GetPlayGround())
		copyController.Snake.len = sc.Snake.len
		copyController.Snake.LastDirection = sc.Snake.LastDirection
		nextStep := sc.GetNextMovableFoodDirection(move)
		nextArray := []DIRECTION{nextStep}
		if Simulate(copyController, nextArray, depth+1, snakeLength) {
			sc.Pg.SetPlayGround(duplicate)
			sc.moveSnakeToFood(nextArray)
		} else {
			sc.Pg.SetPlayGround(duplicate)
			sc.moveSnakeToFood(Remove(move, nextStep))
		}
	} else {
		sc.moveSnakeToFood(move)
	}
	snake := sc.Snake

	if sc.Pg.GetContent(snake.Head.X, snake.Head.Y) == BORDER {
		return false
	} else if sc.Pg.GetContent(snake.Head.X, snake.Head.Y) == FOOD {
		sc.Pg.setRandomFood()
	}
	sc.Pg.DeleteSnake()
	sc.Pg.CreateSnake(snake)

	if snakeLength < 1 {
		return true
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
	for i := 0; i < len(move); i++ {
		//nolint
		if DIRECTION(move[i]) != element {
			//nolint
			movecopy = append(movecopy, DIRECTION(move[i]))
		}
	}
	return movecopy
}
