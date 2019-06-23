package snake

import "testing"

func TestPlayGroundController(t *testing.T) {
	pg := NewPlayground()

	sc := NewSnakeController(pg, NewSnake(5))
	pController := NewPGController(pg, sc, 1)

	if pController == nil {
		t.Errorf("Playground controller is nil")
	}
}

func TestPGControllerNextPeriod(t *testing.T) {
	pg := NewPlayground()
	pg.CreateEmptyPlayground(20, 10)
	pg.CreateOuterBorders()
	snake := NewSnake(5)
	pg.CreateSnake(snake)
	sc := NewSnakeController(pg, snake)
	pController := NewPGController(pg, sc, 1)
	if !pController.nextPeriod(snake) {
		t.Errorf("flag is false; want true")
	}
}

func TestPGController_Start(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("the code did panic")
		}
	}()
	pg := NewPlayground()
	pg.CreateEmptyPlayground(20, 10)
	pg.CreateOuterBorders()
	pg.setStartFood()
	snake := NewSnake(5)
	pg.GetPlayGround()[1][7] = int(BORDER)
	pg.GetPlayGround()[2][6] = int(BORDER)
	pg.CreateSnake(snake)
	sc := NewSnakeController(pg, snake)
	pController := NewPGController(pg, sc, 50)
	pController.Start()
}
