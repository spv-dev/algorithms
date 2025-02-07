package euclid

func Nod(x int, y int) int {
	if x < y {
		x, y = y, x
	}

	if y == 0 {
		return x
	}

	return Nod(y, x%y)
}

func Nod2(x int, y int) int {
	if x < y {
		x, y = y, x
	}

	if y == 0 {
		return x
	}

	if x%y == 0 {
		return y
	}

	return Nod2(y, x%y)
}
