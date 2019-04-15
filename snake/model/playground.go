package model

import (
	"fmt"
)

type Playground interface {
	Create()
}

type playgroundImpl struct {
	playground [][]int
}

func NewPlayground() Playground {
	return &playgroundImpl{
		playground: make([][]int,10),
	}
}

func (pg * playgroundImpl) Create() {
	g := make([][]int, 4)
	counter := 0
	for i := range g {
		g[i] = make([]int, 4)
	}
	for i := range g {
		for j := range g[i] {
			g[i][j] = counter
			counter++
			fmt.Printf("a[%d][%d] = %d\n", i,j, g[i][j] )
		}
	}
	fmt.Printf("%v", g)
}
