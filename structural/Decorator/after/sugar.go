package main

type Sugar struct {
	*CondimentDecorator
}

func NewSugar(beverage Beverage) *Sugar {
	return &Sugar{
		CondimentDecorator: NewCondimentDecorator(beverage),
	}
}

func (s *Sugar) Cost() int {
	return s.beverage.Cost() + 30
}

func (s *Sugar) Description() string {
	return s.beverage.Description() + " + Sugar"
}
