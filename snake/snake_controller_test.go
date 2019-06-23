package snake

import (
	"testing"
)

func TestNewSnakeController(t *testing.T) {
	pg := NewPlayground()
	snake := NewSnake(5)
	sc := NewSnakeController(pg, snake)

	if sc == nil {
		t.Errorf("Snake Controller is nil")
	}
}

func TestSimpleSnakeController_NextStepAllJoices(t *testing.T) {
	pg := NewPlayground()
	pg.CreateEmptyPlayground(20, 10)
	snake := NewSnake(5)
	x, y := snake.Head.X, snake.Head.Y-1
	pg.CreateSnake(snake)
	sc := NewSnakeController(pg, snake)
	sc.NextStep()
	pg.DeleteSnake()
	pg.CreateSnake(sc.GetSnake())
	val := pg.GetContent(x, y)
	// food is not defined, so its on [0][0] --> move up!
	if val != HEAD {
		t.Errorf("val on this field was %d; want %d", int(val), int(HEAD))
	}
}

func TestSimpleSnakeController_NextStepOnlyOneStepPossible(t *testing.T) {
	pg := NewPlayground()
	pg.CreateEmptyPlayground(20, 10)
	pg.CreateOuterBorders()
	pg.setStartFood()
	snake := NewSnake(5)
	pg.CreateSnake(snake)
	pg.GetPlayGround()[1][7] = int(BORDER)
	x, y := snake.Head.X, snake.Head.Y+1
	pg.CreateSnake(snake)
	sc := NewSnakeController(pg, snake)
	sc.NextStep()
	pg.DeleteSnake()
	pg.CreateSnake(sc.GetSnake())
	val := pg.GetContent(x, y)
	//decision is going down
	if val != HEAD {
		t.Errorf("val on this field was %d; want %d", int(val), int(HEAD))
	}
}

/*
he goes left because he would not die, and the way is faster
|------|
|    0 |
|----| |
|$     |
|------|
*/
func TestSimpleSnakeController_RunSimulation(t *testing.T) {
	pg := NewPlayground()
	pg.CreateEmptyPlayground(20, 10)
	pg.CreateOuterBorders()
	pg.setStartFood()
	snake := NewSnake(5)
	snake.Head.X, snake.Head.Y = 5, 1
	tmp := snake.Head
	for i := 5; i >= 1; i-- {
		tmp = tmp.Next
		tmp.X, tmp.Y = i, 2
	}
	x, y := snake.Head.X-1, snake.Head.Y
	pg.CreateSnake(snake)
	sc := NewSnakeController(pg, snake)
	sc.NextStep()
	pg.DeleteSnake()
	pg.CreateSnake(sc.GetSnake())
	val := pg.GetContent(x, y)
	//decision is going down
	if val != HEAD {
		t.Errorf("val on this field was %d; want %d", int(val), int(HEAD))
	}
}

/*
when he goes left, the way is faster but he would die. so he goes right
|------|
| 0    |
|-|    |
||     |
|$     |
|------|
*/
func TestSimpleSnakeController_RunSimulationDontGoIn(t *testing.T) {
	pg := NewPlayground()
	pg.CreateEmptyPlayground(20, 10)
	pg.CreateOuterBorders()
	pg.setStartFood()
	snake := NewSnake(4)
	snake.Head.X, snake.Head.Y = 2, 1
	tmp := snake.Head.Next
	for i := 2; i >= 1; i-- {
		tmp.X, tmp.Y = i, 2
		tmp = tmp.Next
	}
	tmp.X, tmp.Y = 1, 3
	x, y := snake.Head.X+1, snake.Head.Y
	pg.CreateSnake(snake)
	sc := NewSnakeController(pg, snake)
	sc.NextStep()
	pg.DeleteSnake()
	pg.CreateSnake(sc.GetSnake())
	val := pg.GetContent(x, y)
	//decision is going down
	if val != HEAD {
		t.Errorf("val on this field was %d; want %d", int(val), int(HEAD))
	}
}

func TestSimulationRemove(t *testing.T) {
	arr := []DIRECTION{}
	arr = append(arr, UP)
	res := Remove(arr, LEFT)
	if len(res) != 1 {
		t.Errorf("Length is %d; want %d", len(res), 1)
	}
	if res[0] != UP {
		t.Errorf("Direction is %d; want %d", int(res[0]), int(UP))
	}
}
