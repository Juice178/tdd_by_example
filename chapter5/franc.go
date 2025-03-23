package money

type Franc struct {
	amount int
}

func (d *Franc) Times(multiplier int) Franc {
	return Franc{d.amount * multiplier}
}

func (d *Franc) Equals(object Franc) bool {
	return d.amount == object.amount
}
