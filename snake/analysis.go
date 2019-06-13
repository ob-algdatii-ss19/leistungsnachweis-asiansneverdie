package snake

import "fmt"

func GetDirections(playground [][]int) []DIRECTION {
	directions := []DIRECTION{UP, DOWN, LEFT, RIGHT}
	x, y, ok := FindHead(playground)
	if !ok {
		fmt.Println("Head not found")
	}
	// Check UP
	if playground[x][y-1] != int(EMPTY) {
		directions = removeValueFromSlice(directions, UP)
	}
	// check DOWN
	if playground[x][y+1] != int(EMPTY) {
		directions = removeValueFromSlice(directions, DOWN)
	}
	// check LEFT
	if playground[x-1][y] != int(EMPTY) {
		directions = removeValueFromSlice(directions, LEFT)
	}
	// check RIGHT
	if playground[x+1][y] != int(EMPTY) {
		directions = removeValueFromSlice(directions, RIGHT)
	}
	return directions
}

func FindHead(playground [][]int) (int, int, bool) {
	x := -1
	y:= -1
	for i := 0; i < len(playground); i++ {
		for e := 0; e < len(playground[i]); e++ {
			if playground[i][e] == int(HEAD) {
				x = i
				y = e
				return x, y, true
			}
		}
	}
	return x, y, false
}

func removeValueFromSlice(directions []DIRECTION, toRemove DIRECTION) []DIRECTION {
	for i := 0; i < len(directions); i ++ {
		if toRemove == directions[i] {
			directions = append(directions[:i], directions[i+1:]...)
			break
		}
	}
	return directions
}
