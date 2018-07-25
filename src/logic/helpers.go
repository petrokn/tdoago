package logic

func perDelta(start int, end int, delta int) (int, int) {
	current := start

	for current < end && current+delta < end {
		//yield (curr, curr + delta)
		current += delta
		//yield (curr, end)
	}

	return 0, 0
}
