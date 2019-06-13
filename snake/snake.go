package snake

type Snake struct {
	Head *SPart
	len  int
}

type SPart struct {
	X    int
	Y    int
	Next *SPart
}

func NewSnake(len int) Snake {
	snake := Snake{}
	tmp := new(SPart)
	snake.Head = tmp
	snake.len = len
	for i := len + 1; i > 1; i-- {
		tmp.X = i
		tmp.Y = 1
		tmp2 := new(SPart)
		tmp.Next = tmp2
		tmp = tmp2
	}
	tmp.X = 1
	tmp.Y = 1
	return snake
}
