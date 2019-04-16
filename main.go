package main

import "github.com/ob-algdatii-ss19/leistungsnachweis-asiansneverdie/snake"

func main() {
	playground := snake.NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.CreateOuterBorders()
	playground.Print()
}
