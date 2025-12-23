package main

type Whip struct {
	*CondimentDecorator
}

func NewWhip(beverage Beverage) *Whip {
	return &Whip{
		CondimentDecorator: NewCondimentDecorator(beverage),
	}
}

func (w *Whip) Cost() int {
	return w.beverage.Cost() + 70
}

func (w *Whip) Description() string {
	return w.beverage.Description() + " + Whip"
}
