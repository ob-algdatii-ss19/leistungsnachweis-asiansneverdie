package snake

import (
	"testing"
)

func Test_CreatePlayground(t *testing.T) {
	playground := NewPlayground()
	playground.CreateEmptyPlayground(20, 10)

	if playground == nil {
		t.Errorf("playground is nil...")
	}
	if playground.GetPlayGround()[0][0] != int(EMPTY) {
		t.Errorf("first field on playground is: %d; want %d", playground.GetPlayGround()[0][0], int(EMPTY))
	}
}

func TestPlaygroundImpl_GetPlayGround(t *testing.T) {
	playground := NewPlayground()
	playground.CreateEmptyPlayground(20, 10)

	val := playground.GetPlayGround()
	if val == nil {
		t.Error("Playground is nil...")
	}
}

func TestPlaygroundImpl_CreateOuterBorders(t *testing.T) {
	playground := NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.CreateOuterBorders()

	val := playground.GetPlayGround()
	if val[0][0] != int(BORDER) {
		t.Errorf("0/0 is no Border")
	}
}

func TestPlaygroundImpl_CreateSnake(t *testing.T) {
	playground := NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	s := NewSnake(5)
	playground.CreateSnake(s)
	val := playground.GetPlayGround()
	tmp := s.Head
	if val[tmp.Y][tmp.X] != int(HEAD) {
		t.Errorf("head is not at coords %d-%d", tmp.X, tmp.Y)
	}
	tmp = tmp.Next
	for tmp != nil {
		if val[tmp.Y][tmp.X] != int(TAIL) {
			t.Errorf("tail is not at coords %d-%d", tmp.X, tmp.Y)
		}
		tmp = tmp.Next
	}
}

func TestPlaygroundImpl_SetPlayGround(t *testing.T) {
	playground := NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	val := make([][]int, 10)
	for i := range val {
		val[i] = make([]int, 20)
	}
	val[0][0] = int(HEAD)

	playground.SetPlayGround(val)
	if playground.GetPlayGround()[0][0] != val[0][0] {
		t.Errorf("field is %d; want %d", playground.GetPlayGround()[0][0], val[0][0])
	}
}

func TestPlaygroundImpl_CopyPlayGround(t *testing.T) {
	playground := NewPlayground()
	val := make([][]int, 10)
	for i := range val {
		val[i] = make([]int, 20)
	}
	val[0][0] = int(HEAD)
	other := playground.CopyPlayGround(val)
	if other[0][0] != int(HEAD) {
		t.Errorf("first field is %d; want %d", other[0][0], val[0][0])
	}
}

func TestPlaygroundImpl_DeleteSnake(t *testing.T) {
	playground := NewPlayground()
	snake := NewSnake(5)
	playground.CreateEmptyPlayground(20, 10)
	playground.CreateSnake(snake)
	playground.DeleteSnake()
	val := playground.GetPlayGround()
	tmp := snake.Head
	if val[tmp.Y][tmp.X] == int(HEAD) {
		t.Errorf("head is still at coords %d-%d", tmp.X, tmp.Y)
	}
	tmp = tmp.Next
	for tmp != nil {
		if val[tmp.Y][tmp.X] == int(TAIL) {
			t.Errorf("tail is still at coords %d-%d", tmp.X, tmp.Y)
		}
		tmp = tmp.Next
	}
}

func TestPlayGroundImpls_createRandomFood(t *testing.T) {
	playground := NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.setRandomFood()
	suc := false
	for i := range playground.GetPlayGround() {
		for j := range playground.GetPlayGround()[i] {
			if playground.GetPlayGround()[i][j] == int(FOOD) {
				suc = true
			}
		}
	}
	if !suc {
		t.Errorf("no food was found")
	}
}

func TestPlaygroundImpl_GetFood(t *testing.T) {
	playground := NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.setRandomFood()
	x, y := playground.GetFood()
	if playground.GetPlayGround()[y][x] != int(FOOD) {
		t.Errorf("no food was found")
	}
}

func TestPlaygroundImpl_GetContent(t *testing.T) {
	playground := NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.CreateOuterBorders()
	con := playground.GetContent(0, 0)
	if con != BORDER {
		t.Errorf("content is %d; want %d", int(con), int(BORDER))
	}
}



