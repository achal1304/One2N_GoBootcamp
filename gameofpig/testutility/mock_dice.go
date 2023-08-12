package testutility

type TestDice struct {
	values []int
	index  int
}

func NewTestDice(values []int) *TestDice {
	return &TestDice{
		values: values,
		index:  0,
	}
}

func (td *TestDice) Roll() int {
	result := td.values[td.index]
	td.index = (td.index + 1) % len(td.values)
	return result
}
