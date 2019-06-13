package snake

type DIRECTION int

const (
	UP    DIRECTION = 0
	DOWN  DIRECTION = 1
	LEFT  DIRECTION = 2
	RIGHT DIRECTION = 3
)

type SimpleSnakeController struct {
	pg    Playground
	Snake Snake
}

func NewSnakeController(pg Playground, snake Snake) SController {
	sc := new(SimpleSnakeController)
	sc.pg = pg
	sc.Snake = snake

	return sc
}

func (sc *SimpleSnakeController) NextStep() {
	//Simple Implementatoin - Just move one step forwarads
	//var x, y = sc.pg.GetFood()


	sc.moveSnake()
}

func (sc *SimpleSnakeController) GetSnake() Snake {
	return sc.Snake
}

func (sc *SimpleSnakeController) moveSnake() {
	dir := DOWN
	var x, y = sc.pg.GetFood()
	if sc.Snake.Head.x < x {
		dir = RIGHT
	} else if sc.Snake.Head.x > x {
		dir = LEFT
	} else if sc.Snake.Head.y < y {
		dir = DOWN
	} else {
		dir = UP
	}
	// if snake got food dont delete the last tail
	if sc.getNextPGField(dir) != FOOD {
		sc.setLastTail()
	}
	sc.setNewHead(dir)
}

func (sc *SimpleSnakeController) addTail() Snake {
	sc.Snake = NewSnake(sc.GetSnake().len + 1)
	return sc.Snake
}

//set the second last tail to nil
func (sc *SimpleSnakeController) setLastTail() {
	first := sc.Snake.Head
	second := first.next
	for {
		if second.next == nil {
			break
		}
		first = second
		second = second.next
	}
	first.next = nil
}

func (sc *SimpleSnakeController) setNewHead(dir DIRECTION) {
	newHead := new(SPart)
	newHead.x, newHead.y = sc.getNextSnakeField(dir)
	newHead.next = sc.Snake.Head
	sc.Snake.Head = newHead
}

func (sc *SimpleSnakeController) getNextPGField(dir DIRECTION) CONTENT {
	x, y := sc.getNextSnakeField(dir)
	return sc.pg.GetContent(x, y)
}

func (sc *SimpleSnakeController) getNextSnakeField(dir DIRECTION) (int, int) {
	s := sc.GetSnake().Head
	x, y := s.x, s.y
	switch dir {
	case UP:
		y = y-1
		sc.Snake.LastDirection = UP
	case DOWN:
		y = y+1
		sc.Snake.LastDirection = DOWN
	case RIGHT:
		x = x+1
		sc.Snake.LastDirection = RIGHT
	case LEFT:
		x = x-1
		sc.Snake.LastDirection = LEFT
	default:
		// do nothing
	}
	return x, y
}
