package command

import "fmt"

type Order interface {
	execute() error
}

//
type Stock struct {
	name 		string
	quantity 	int
}

func NewStock(name string, quantity int) *Stock {
	return &Stock{
		name: name,
		quantity: quantity,
	}
}

func (s *Stock) Buy() {
	fmt.Printf("Stock [Name: %s, Quantity: %d] bought.\n", s.name, s.quantity)
	return
}

func (s *Stock) Sell() {
	fmt.Printf("Stock [Name: %s, Quantity: %d] sold.\n", s.name, s.quantity)
	return
}

// implement Order
type BuyStock struct {
	stock 	*Stock
}

func NewBuyStock(stock *Stock) Order {
	return &BuyStock{
		stock: stock,
	}
}

func (bs *BuyStock) execute() error {
	bs.stock.Buy()
	return nil
}

type SellStock struct {
	stock 	*Stock
}

func NewSellStock(stock *Stock) Order {
	return &SellStock{
		stock: stock,
	}
}

func (ss *SellStock) execute() error {
	ss.stock.Sell()
	return nil
}

// broker
type Broker struct {
	OrderList 	[]Order
}

func NewBroker() *Broker {
	return &Broker{
		OrderList: make([]Order, 0),
	}
}

func (b *Broker) TakeOrder(order Order) {
	b.OrderList = append(b.OrderList, order)
}

func (b *Broker) PlaceOrders() {
	for _, order := range b.OrderList {
		order.execute()
	}
	b.OrderList = b.OrderList[:0]
}
