package money

import "fmt"

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
