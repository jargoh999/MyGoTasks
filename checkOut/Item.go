package main

type Item struct {
	id               int
	price            float64
	products         []Product
	numberOfProducts int
}

func (item *Item) addProducts(product Product) {
	item.products = append(item.products, product)
	item.numberOfProducts++
}

func (item *Item) removeProducts(productToRemove Product) {
	for product := 0; product < item.numberOfProducts; product++ {
		if item.products[product].Id == productToRemove.getId() {
			item.products = append(item.products[:product], item.products[product+1:]...)
			item.numberOfProducts--
		}
	}
	return
}

func (item *Item) setId(id int) {
	item.id = id
}

func (item *Item) getId() int {
	return item.id
}

func (item *Item) getNumberOfProducts() int {
	return item.numberOfProducts
}

func (item *Item) getPrice() float64 {
	totalPrice := 0.0
	for _, product := range item.products {
		totalPrice += product.price
	}
	return totalPrice
}
