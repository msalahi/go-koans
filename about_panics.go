package go_koans

func divideFourBy(i int) int {
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	assert(true) // panics are exceptional errors at runtime

	n := divideFourBy(2)
	assert(n == 2) // panics are exceptional errors at runtime
}
