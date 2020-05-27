package go_koans

func divideFourBy(i int) int {
	return 4 / i
}

const __divisor__ = 0
const __properDivisor__ = 2

func aboutPanics() {
	//assert(__delete_me__) // panics are exceptional errors at runtime
	defer func() {
		if r := recover(); r != nil {
			if  err, ok := r.(error); ok && err != nil {
				n := divideFourBy(__properDivisor__)
				assert(n == 2)
			}
		}
	}()
	n := divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
