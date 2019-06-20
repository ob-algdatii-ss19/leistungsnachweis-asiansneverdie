package snake

// TODO: correct input
func Simulate(sc *SimpleSnakeController, move []DIRECTION, snakeLength int) bool {
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
}
