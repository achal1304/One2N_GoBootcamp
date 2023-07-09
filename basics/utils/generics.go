package utils

func All[T any](conditions []func(T) bool, value T) bool {
	for _, f := range conditions {
		if !f(value) {
			return false
		}
	}
	return true
}

func Any[T any](conditions []func(T) bool, value T) bool {
	for _, f := range conditions {
		if f(value) {
			return true
		}
	}
	return false
}
