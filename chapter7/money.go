package money

import "fmt"

type Money interface {
	Equals(object any) bool
}

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

type Franc struct {
	amount int
}

func (d *Franc) Times(multiplier int) Franc {
	return Franc{d.amount * multiplier}
}

func (d *Franc) Equals(object any) bool {
	franc, ok := object.(Franc)
	if !ok {
		fmt.Printf("Argument is no type of Dollar")
		return false
	}
	return d.amount == franc.amount
}
