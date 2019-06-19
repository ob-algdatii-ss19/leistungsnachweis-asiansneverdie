package snake

func simulate(pg [][]int, len int) bool {

	// TODO: simulate here

	len = len - 1
	if len <= 0 {
		return true
	} else {
		return simulate(pg, len)
	}
}
