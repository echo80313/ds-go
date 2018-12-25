package ds

// LeastSignificantBit clears all bits except the
// least significant one.
func LeastSignificantBit(x int) int {
	return x & (-x)
}

// NextPowerOf2 find the next power of 2 greater or
// equal to x.
func NextPowerOf2(x int) int {
	t := 1
	for i := 0; i < 31; i++ {
		if t >= x {
			return t
		}
		t = 2 * t
	}
	return t
}
