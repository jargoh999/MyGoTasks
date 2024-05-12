package main

type Store struct {
	products []Product
}

func (store *Store) getAllProducts() []Product {
	return store.products
}

func (store *Store) addProduct(admin Admin) {
	store.products = append(store.products, admin.getProduct()...)
}

func (store *Store) removeProduct(id int) {
	for product, adminProduct := range store.products {
		if adminProduct.Id == id {
			store.products = append(store.products[:product], store.products[product+1:]...)
		}
	}
	return
}

func (store *Store) getProductById(id int) Product {
	for product, adminProduct := range store.products {
		if adminProduct.Id == id {
			return store.products[product]
		}
	}
	return Product{}
}
