package ch4

func AppendInt(x []int, y int) []int {
	var z []int
	zLen := len(x) + 1

	if zLen <= cap(x) {
		// There's room to grow. Extend the slice
		z = x[:zLen]
	} else {
		// There is insufficient space. Allotate a new array/
		// Grow by doubling, for amortized linear complexity
		zcap := zLen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zLen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}
