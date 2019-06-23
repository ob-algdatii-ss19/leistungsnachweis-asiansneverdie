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
