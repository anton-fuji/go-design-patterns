package main

// トッピングのデコレータ
type CondimentDecorator struct {
	beverage Beverage
}

func NewCondimentDecorator(beverage Beverage) *CondimentDecorator {
	return &CondimentDecorator{
		beverage: beverage,
	}
}

func (c *CondimentDecorator) Cost() int {
	return c.beverage.Cost()
}

func (c *CondimentDecorator) Description() string {
	return c.beverage.Description()
}
