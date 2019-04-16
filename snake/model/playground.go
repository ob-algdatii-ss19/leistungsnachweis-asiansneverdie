package model

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
	Print()
}

type playgroundImpl struct {
	playground [][]int
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

func (pg *playgroundImpl) Print() {
	for i := range pg.playground {
		for j := range pg.playground[i] {
			switch pg.playground[i][j] {
			case int(EMPTY):
				print("  ")
			case int(BORDER):
				print(" #")
			case int(HEAD):
				print(" o")
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
