package snake

func Simulate(pg Playground, len int, s Snake) bool {

	// TODO: simulate here


	len = len - 1
	if len <= 0 {
		return true
	} else {
		return Simulate(pg, len, s)
	}
}
