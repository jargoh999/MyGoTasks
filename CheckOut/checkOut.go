package main

type Item struct {
	id               int
	Price            float64
	products         []Product
	numberOfProducts int
}

func (item *Item) addProducts(product Product) {
	item.products = append(item.products, product)
	item.numberOfProducts++
}

func (item *Item) removeProducts(productId int) {
	for product := 0; product < item.numberOfProducts; product++ {
		if item.products[product].Id == productId {
			item.products = append(item.products[:product], item.products[product+1:]...)
			item.numberOfProducts--
		}
	}
	return
}

func (item *Item) getNumberOfProducts() int {
	return item.numberOfProducts
}

func (item *Item) getProductPrice() float64 {
	totalPrice := 0.0
	for _, product := range item.products {
		totalPrice += product.price
	}
	return totalPrice
}

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
	return 0
}

func (customer *Customer) setName(name string) {
	customer.name = name
}

type Store struct {
	products []Product
}

func (store *Store) getAllProducts() []Product {
	return store.products
}

func (store *Store) addProduct(admin Admin) {
	store.products = append(store.products, admin.getProduct()...)
}

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

type Cart struct {
	id    int
	items []Item
}

func (cart *Cart) addItem(item Item) {
	cart.items = append(cart.items, item)
}

func (cart *Cart) removeItem(itemId int) {
	for item := 0; item < len(cart.items); item++ {
		if cart.items[item].id == itemId {
			cart.items = append(cart.items[:item], cart.items[item+1:]...)
		}
	}
	return
}

func (cart *Cart) getNumberOfItems() int {
	return len(cart.items)
}

func (cart *Cart) getItems() []Item {
	return cart.items
}

type Admin struct {
	id       int
	name     string
	products []Product
}

func (admin *Admin) getProduct() []Product {
	return admin.products
}

func (admin *Admin) addProduct(product Product) {
	admin.products = append(admin.products, product)
}

func (admin *Admin) setId(id int) {
	admin.id = id
}

func (admin *Admin) getId() int {
	return admin.id
}

func (admin *Admin) setName(name string) {
	admin.name = name
}

func (admin *Admin) getName() string {
	return admin.name
}
