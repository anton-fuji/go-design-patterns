package main

type Coffee struct {
	name string
}

func NewCoffee(name string) *Coffee {
	return &Coffee{
		name: name,
	}
}

func (c *Coffee) Cost() int {
	return 300
}

func (c *Coffee) Description() string {
	return c.name
}
