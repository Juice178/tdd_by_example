package main

func main() {
}

type Dollar struct {
	amount int
}

func (d *Dollar) Times(multiplier int) {
	d.amount *= multiplier
}
