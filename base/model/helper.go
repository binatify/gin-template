package model

func EnsureWithinMaxItems(total int) int {
	if total > DefaultMaxItems || total < DefaultMinItems {
		total = DefaultMaxItems
	}

	return total
}
