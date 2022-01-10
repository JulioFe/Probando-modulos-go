package course

import "fmt"

type course struct {
	name    string
	UsersID []uint
	Clasess map[uint]string
	Price   uint
}

func (c *course) setName(name string) {
	c.name = name
}

func (c *course) Name() string {
	return c.name
}

func New(name string, price uint) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		name:  name,
		Price: price,
	}
}

func (c *course) AllClasess() {
	text := ""
	for _, v := range c.Clasess {
		text += v + ","
	}
	fmt.Println(text[:len(text)-1])
}

func (c *course) SetPrice(new_price uint) {
	c.Price = new_price
}
