package snake

import "testing"

func TestSnake(t *testing.T) {
	snakeL := 5
	snake := NewSnake(snakeL)

	if snake.len != 5 {
		t.Errorf("len of snake is: %d; want %d", snake.len, snakeL)
	}
	if snake.Head.X != snakeL+1 {
		t.Errorf("x coords of snake head is: %d; want %d", snake.Head.X, snakeL+1)
	}
}
