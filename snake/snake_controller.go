package snake

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/mohae/deepcopy"
)

type DIRECTION int

const (
	UP    DIRECTION = 0
	DOWN  DIRECTION = 1
	LEFT  DIRECTION = 2
	RIGHT DIRECTION = 3
)

type SimpleSnakeController struct {
	Pg    Playground
	Snake Snake
}

func NewSnakeController(pg Playground, snake Snake) SController {
	sc := new(SimpleSnakeController)
	sc.Pg = pg
	sc.Snake = snake

	return sc
}

func (sc *SimpleSnakeController) NextStep() {
	move := GetDirections(sc.Pg.GetPlayGround())
	switch len(move) {
	case 0:
		// Default, End Game
		sc.setNewHead(RIGHT)
	case 1:
		// move in the one direction
		sc.setLastTail()
		sc.setNewHead(move[0])
	case 2:
		copied := deepcopy.Copy(sc)
		copyController := copied.(*SimpleSnakeController)
		err := copier.Copy(copyController.Pg, sc.Pg)
		if err != nil {
			fmt.Println("Error occurred while copying")
		}
		duplicate := copyController.Pg.CopyPlayGround(copyController.Pg.GetPlayGround())
		copyController.Snake.len = sc.Snake.len
		copyController.Snake.LastDirection = sc.Snake.LastDirection
		nextStep := copyController.GetNextMovableFoodDirection(move)
		nextArray := []DIRECTION{nextStep}
		if Simulate(copyController, nextArray, 0, copyController.Snake.len) {
			sc.Pg.SetPlayGround(duplicate)
			sc.moveSnakeToFood(nextArray)
		} else {
			sc.Pg.SetPlayGround(duplicate)
			sc.moveSnakeToFood(Remove(move, nextStep))
		}
	case 3:
		// move to food
		sc.moveSnakeToFood(move)
	default:
	}
}

func (sc *SimpleSnakeController) GetSnake() Snake {
	return sc.Snake
}

func (sc *SimpleSnakeController) moveSnakeToFood(move []DIRECTION) {
	dir := move[0]
	var x, y = sc.Pg.GetFood()
	if sc.Snake.Head.X < x && contains(move, RIGHT) {
		dir = RIGHT
	} else if sc.Snake.Head.X > x && contains(move, LEFT) {
		dir = LEFT
	} else if sc.Snake.Head.Y < y && contains(move, DOWN) {
		dir = DOWN
	} else if contains(move, UP) {
		dir = UP
	}
	// if snake got food dont delete the last tail
	if sc.getNextPGField(dir) != FOOD {
		sc.setLastTail()
	} else {
		sc.Snake.len++
	}
	sc.setNewHead(dir)
}

func contains(s []DIRECTION, e DIRECTION) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (sc *SimpleSnakeController) addTail() Snake {
	sc.Snake = NewSnake(sc.GetSnake().len + 1)
	return sc.Snake
}

//set the second last tail to nil
func (sc *SimpleSnakeController) setLastTail() {
	first := sc.Snake.Head
	second := first.Next
	for {
		if second.Next == nil {
			break
		}
		first = second
		second = second.Next
	}
	first.Next = nil
}

func (sc *SimpleSnakeController) setNewHead(dir DIRECTION) {
	newHead := new(SPart)
	newHead.X, newHead.Y = sc.getNextSnakeField(dir)
	newHead.Next = sc.Snake.Head
	sc.Snake.Head = newHead
}

func (sc *SimpleSnakeController) getNextPGField(dir DIRECTION) CONTENT {
	x, y := sc.getNextSnakeField(dir)
	return sc.Pg.GetContent(x, y)
}

func (sc *SimpleSnakeController) getNextSnakeField(dir DIRECTION) (int, int) {
	s := sc.GetSnake().Head
	x, y := s.X, s.Y
	switch dir {
	case UP:
		y = y - 1
		sc.Snake.LastDirection = UP
	case DOWN:
		y = y + 1
		sc.Snake.LastDirection = DOWN
	case RIGHT:
		x = x + 1
		sc.Snake.LastDirection = RIGHT
	case LEFT:
		x = x - 1
		sc.Snake.LastDirection = LEFT
	default:
		// do nothing
	}
	return x, y
}
