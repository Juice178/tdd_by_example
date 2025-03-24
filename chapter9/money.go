package money

import (
	"fmt"
)

type Money interface {
	Equals(object any) bool
	Times(multiplier int) Money
	Currency() string
}

type Dollar struct {
	amount   int
	currency string
}

func (d *Dollar) Times(multiplier int) Money {
	return &Dollar{d.amount * multiplier, "USD"}
}

func (d *Dollar) Equals(object any) bool {
	dollar, ok := object.(Dollar)
	if !ok {
		fmt.Printf("Argument is not type of Dollar")
		return false
	}
	return d.amount == dollar.amount
}

func (d *Dollar) Currency() string {
	return d.currency
}

func newDollar(amount int) Money {
	return &Dollar{amount: amount, currency: "USD"}
}

type Franc struct {
	amount   int
	currency string
}

func (d *Franc) Times(multiplier int) Money {
	return &Franc{d.amount * multiplier, "CHF"}
}

func (d *Franc) Equals(object any) bool {
	franc, ok := object.(Franc)
	if !ok {
		fmt.Printf("Argument is not type of Franc")
		return false
	}
	return d.amount == franc.amount
}

func (d *Franc) Currency() string {
	return d.currency
}

func newFranc(amount int) Money {
	return &Franc{amount: amount, currency: "CHF"}
}

func GetMoney(currency string, amount int) (Money, error) {
	if currency == "Dollar" {
		return newDollar(amount), nil
	}
	if currency == "Franc" {
		return newFranc(amount), nil
	}
	return nil, fmt.Errorf("wrong currency type passed")
}
