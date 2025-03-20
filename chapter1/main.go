package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

type Dollar struct {
	amount int
}

func (d *Dollar) Times(multiplier int) {
	d.amount = 10
}
