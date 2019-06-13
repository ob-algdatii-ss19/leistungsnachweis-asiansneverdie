package snake

import (
	"testing"
)

func TestSnake_CreatePlayground(t *testing.T) {
	playground := NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.CreateOuterBorders()
	snake := NewSnake(4)
	playground.CreateSnake(snake)
	playground.Print()
}
