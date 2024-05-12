package main

type Customer struct {
	id      int
	cart    Cart
	name    string
	balance float64
}

func (customer *Customer) getCart() Cart {
	return customer.cart
}

func (customer *Customer) getBalance(cart Cart) float64 {
	return cart.getPriceOfItems()
}

func (customer *Customer) setName(name string) {
	customer.name = name
}
func (customer *Customer) getName() string {
	return customer.name
}
func (Customer *Customer) setId(id int) {
	Customer.id = id
}
func (Customer *Customer) getId() int {
	return Customer.id
}
