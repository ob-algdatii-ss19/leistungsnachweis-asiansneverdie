package controller

import "github.com/ob-algdatii-ss19/leistungsnachweis-asiansneverdie/snake/model"

func Main() {
	playground := model.NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.CreateOuterBorders()
	playground.Print()
}
