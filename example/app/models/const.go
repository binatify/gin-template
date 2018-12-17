package models

const (
	DefaultMaxItems = 100
	DefaultMinItems = 1
)

func EnsureWithinMaxItems(total int) int {
	if total > DefaultMaxItems || total < DefaultMinItems {
		total = DefaultMaxItems

	}

	return total
}
