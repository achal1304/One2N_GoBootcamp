package testutility

type DiceStub struct {
	values []int
	index  int
}

func NewTestDice(values []int) *DiceStub {
	return &DiceStub{
		values: values,
		index:  0,
	}
}

func (td *DiceStub) Roll() int {
	result := td.values[td.index]
	td.index = (td.index + 1) % len(td.values)
	return result
}
