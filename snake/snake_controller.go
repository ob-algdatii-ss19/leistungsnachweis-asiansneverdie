package snake

type DIRECTION int

const (
	UP    DIRECTION = 0
	DOWN  DIRECTION = 1
	LEFT  DIRECTION = 2
	RIGHT DIRECTION = 3
)

type SimpleSnakeController struct {
	pg    *Playground
	Snake Snake
}

func NewSnakeController(pg *Playground, snake Snake) SController {
	sc := new(SimpleSnakeController)
	sc.pg = pg
	sc.Snake = snake

	return sc
}

func (sc *SimpleSnakeController) NextStep() {
	//Simple Implementatoin - Just move one step forwarads
	sc.moveSnake()
}

func (sc *SimpleSnakeController) GetSnake() Snake {
	return sc.Snake
}

func (sc *SimpleSnakeController) moveSnake() {
	sc.setLastTail()
	sc.setNewHead(RIGHT)
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
	switch dir {
	case UP:
		newHead.x = sc.Snake.Head.x
		newHead.y = sc.Snake.Head.y - 1
	case DOWN:
		newHead.x = sc.Snake.Head.x
		newHead.y = sc.Snake.Head.y + 1
	case LEFT:
		newHead.x = sc.Snake.Head.x - 1
		newHead.y = sc.Snake.Head.y
	case RIGHT:
		newHead.x = sc.Snake.Head.x + 1
		newHead.y = sc.Snake.Head.y
	default:
		// do nothing
	}
	newHead.next = sc.Snake.Head
	sc.Snake.Head = newHead
}
