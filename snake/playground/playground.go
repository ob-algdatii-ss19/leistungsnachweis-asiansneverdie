package playground

import (
	"math/rand"
	"time"
)

type CONTENT int

const (
	EMPTY  CONTENT = 0
	BORDER CONTENT = 1
	HEAD   CONTENT = 2
	TAIL   CONTENT = 3
	FOOD   CONTENT = 4
)


type Playground interface {
	CreateEmptyPlayground(height, width int)
	CreateOuterBorders()
	CreateSnake(snake Snake)
	setStartFood()
	setRandomFood()
	Print()
	DeleteSnake()
	GetContent(x, y int) CONTENT
	GetFood() (int, int)

}

type playgroundImpl struct {
	playground [][]int
	foodX int
	foodY int
}

func NewPlayground() Playground {
	return &playgroundImpl{
		playground: nil,
	}
}

func (pg *playgroundImpl) CreateEmptyPlayground(height, width int) {
	g := make([][]int, width)
	for i := range g {
		g[i] = make([]int, height)
	}
	for i := range g {
		for j := range g[i] {
			g[i][j] = int(EMPTY)
		}
	}
	pg.playground = g
}

func (pg *playgroundImpl) CreateOuterBorders() {
	for i := range pg.playground {
		for j := range pg.playground[i] {
			if i == 0 || i == len(pg.playground)-1 {
				pg.playground[i][j] = int(BORDER)
			}
			if j == 0 || j == len(pg.playground[0])-1 {
				pg.playground[i][j] = int(BORDER)
			}
		}
	}
}

func (pg *playgroundImpl) CreateSnake(snake Snake) {
	part := snake.Head
	// Draw Head
	pg.playground[part.y][part.x] = int(HEAD)
	// Draw Tail
	for {
		part = part.next
		if part == nil {
			break
		}
		pg.playground[part.y][part.x] = int(TAIL)
	}
}

func (pg *playgroundImpl) DeleteSnake() {
	for i := range pg.playground {
		for j := range pg.playground[i] {
			tmp := pg.playground[i][j]
			if CONTENT(tmp) == HEAD || CONTENT(tmp) == TAIL {
				pg.playground[i][j] = int(EMPTY)
			}
		}
	}
}

func (pg *playgroundImpl) GetContent(x, y int) CONTENT {
	return CONTENT(pg.playground[y][x])
}

func (pg *playgroundImpl) Print() {
	for i := range pg.playground {
		for j := range pg.playground[i] {
			switch pg.playground[i][j] {
			case int(EMPTY):
				print("  ")
			case int(BORDER):
				print(" #")
			case int(HEAD):
				print(" O")
			case int(TAIL):
				print(" +")
			case int(FOOD):
				print(" $")
			default:
				// just for linter
			}
		}
		print("\n")
	}
	print("\n")
}

func (pg *playgroundImpl) setRandomFood() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(len(pg.playground) - 3) + 1
	y := rand.Intn(len(pg.playground[0]) - 3) + 1
	for {
		if CONTENT(pg.playground[y][x]) == EMPTY {
			pg.playground[y][x] = int(FOOD)
			pg.foodX = x
			pg.foodY = y
			return
		}
	}
}

func (pg *playgroundImpl) GetFood() (int, int) {
	return pg.foodX, pg.foodY
}

func (pg *playgroundImpl) setStartFood() {
	pg.playground[5][5] = int(FOOD);
}
