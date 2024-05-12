package main

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

func (cart *Cart) getPriceOfItems() float64 {
	totalPrice := 0.0
	for _, item := range cart.items {
		totalPrice += item.getPrice()
	}
	return totalPrice
}

func (cart *Cart) getId() int {
	return cart.id
}

func (cart *Cart) setId(id int) {
	cart.id = id
}
