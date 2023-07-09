package story7

func MultiConditionsCheck[T any](values []T, conditionCheck func([]func(T) bool, T) bool, conditions []func(T) bool) []T {
	result := []T{}

	for _, v := range values {
		if conditionCheck(conditions, v) {
			result = append(result, v)
		}
	}
	return result
}
