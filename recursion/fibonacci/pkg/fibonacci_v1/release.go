package fibonacci

func Find(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return Find(n-1) + Find(n-2)
}
