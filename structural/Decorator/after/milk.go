package main

// ミルクトッピング
type Milk struct {
	*CondimentDecorator
}

func NewMilk(beverage Beverage) *Milk {
	return &Milk{
		CondimentDecorator: NewCondimentDecorator(beverage),
	}
}

func (m *Milk) Cost() int {
	return m.beverage.Cost() + 50
}

func (m *Milk) Description() string {
	return m.beverage.Description() + " + Milk"
}
