package playground

import "time"

type gameController struct {
	pg    Playground
	sc    SController
	speed int
}

// returns a new PlayGroundController
func NewPGController(pg Playground, sc SController, speed int) PlayGroundController {
	gc := new(gameController)
	gc.pg = pg
	gc.sc = sc
	gc.speed = speed
	return gc
}

// Starts the Snake
func (gc *gameController) Start() {
	snake := gc.sc.GetSnake()
	gc.pg.setStartFood()
	for {
		gc.pg.Print()
		time.Sleep(time.Duration(gc.speed) * time.Millisecond)
		gc.sc.NextStep()
		snake = gc.sc.GetSnake()
		if gc.pg.GetContent(snake.Head.x, snake.Head.y) == BORDER {
			break
		} else if gc.pg.GetContent(snake.Head.x, snake.Head.y) == FOOD {
			gc.pg.setRandomFood()
		}
		gc.pg.DeleteSnake()
		gc.pg.CreateSnake(snake)
	}
}
