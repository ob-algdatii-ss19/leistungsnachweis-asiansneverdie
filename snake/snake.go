package snake

type Snake struct {
	Head *SPart
	len  int
}

type SPart struct {
	x    int
	y    int
	next *SPart
}

func NewSnake(len int) Snake {
	snake := Snake{}
	tmp := new(SPart)
	snake.Head = tmp
	snake.len = len

	for i := len + 1; i > 1; i-- {
		tmp.x = 1
		tmp.y = i
		tmp2 := new(SPart)
		tmp.next = tmp2
		tmp = tmp2
	}
	tmp.x = 1
	tmp.y = 1
	return snake
}
