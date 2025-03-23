package money

import "fmt"

type Dollar struct {
	amount int
}

func (d *Dollar) Times(multiplier int) Dollar {
	return Dollar{d.amount * multiplier}
}

func (d *Dollar) Equals(object any) bool {
	dollar, ok := object.(Dollar)
	if !ok {
		fmt.Printf("Argument is no type of Dollar")
		return false
	}
	return d.amount == dollar.amount
}
