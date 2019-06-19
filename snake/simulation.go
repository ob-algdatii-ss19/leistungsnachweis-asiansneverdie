package snake

// TODO: correct input
func Simulate(pg Playground, len int, s Snake) bool {

	// TODO: simulate here
	// TODO: COPY ALL REFERENCES
	// TODO: SIMULATE GO TO FOOD

	len = len - 1
	if len <= 0 {
		return true
	} else {
		return Simulate(pg, len, s)
	}
}
