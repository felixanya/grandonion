package command

import "testing"

func TestNewBroker(t *testing.T) {
	abcStock := NewStock("abc", 500)
	defStock := NewStock("def", 200)

	buy := NewBuyStock(abcStock)
	sell := NewSellStock(defStock)

	broker := NewBroker()
	broker.TakeOrder(buy)
	broker.TakeOrder(sell)

	broker.PlaceOrders()
}
