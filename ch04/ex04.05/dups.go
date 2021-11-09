package ch04

func FilterDups(items []string) []string {
	l := len(items)
	if l == 0 {
		return items
	}

	n := 1
	for i := 1; i < l; i++ {
		if items[i] != items[i-1] {
			items[n] = items[i]
			n++
		}
	}
	items = items[:n]

	return items
}
