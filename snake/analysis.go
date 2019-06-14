package snake

import "fmt"

func GetDirections(playground [][]int) []DIRECTION {
	directions := []DIRECTION{UP, DOWN, LEFT, RIGHT}
	x, y, ok := FindHead(playground)
	if !ok {
		fmt.Println("Head not found")
	}
	// Check UP
	if playground[y-1][x] != int(EMPTY) && playground[y-1][x] != int(FOOD) {
		directions = removeValueFromSlice(directions, UP)
	}
	// check DOWN
	if playground[y+1][x] != int(EMPTY) &&  playground[y+1][x] != int(FOOD) {
		directions = removeValueFromSlice(directions, DOWN)
	}
	// check LEFT
	if playground[y][x-1] != int(EMPTY) && playground[y][x-1] != int(FOOD){
		directions = removeValueFromSlice(directions, LEFT)
	}
	// check RIGHT
	if playground[y][x+1] != int(EMPTY) && playground[y][x+1] != int(FOOD) {
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
				x = e
				y = i
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
