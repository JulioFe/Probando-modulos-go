package articulo

type articulo struct {
	name  string
	price uint16
}

func New(name string, price uint16) articulo {
	return articulo{name, price}
}

type articulos []articulo

func (art articulos) GetSum() int {
	sum := 0
	for _, v := range art {
		sum += int(v.price)
	}
	return sum
}
