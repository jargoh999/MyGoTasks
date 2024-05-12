package main

type Product struct {
	Id    int
	name  string
	price float64
}

func (product *Product) getId() int {
	return product.Id
}

func (product *Product) setId(id int) {
	product.Id = id
}

func (product *Product) setName(name string) {
	product.name = name
}

func (product *Product) getName() string {
	return product.name
}

func (product *Product) setPrice(price float64) {
	product.price = price
}
