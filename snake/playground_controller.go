package snake

func Main() {
	playground := NewPlayground()
	playground.CreateEmptyPlayground(20, 10)
	playground.CreateOuterBorders()
	playground.Print()
}
