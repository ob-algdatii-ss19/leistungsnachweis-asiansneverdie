package controller

import "github.com/Plateria/AlgoDatII/leistungsnachweis-asiansneverdie/snake/model"

func Main() {
	playground := model.NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.CreateOuterBorders()
	playground.Print()
}
