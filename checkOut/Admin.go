package main

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
