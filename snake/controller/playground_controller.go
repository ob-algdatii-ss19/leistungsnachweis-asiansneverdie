package controller

import "../model"

func Main() {
	playground := model.NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.CreateOuterBorders()
	playground.Print()
}
