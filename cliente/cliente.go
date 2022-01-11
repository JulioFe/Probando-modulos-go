package cliente

type cliente struct {
	name    string
	age     uint8
	partner bool
}

func New(name string, age uint8, partner bool) cliente {
	return cliente{name, age, partner}
}
