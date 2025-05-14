package entity

import (
	"sync"

	//"golang.org/x/tools/go/analysis/passes/defers"
)

type Book struct {
	Order           []*Order
	Transactions    []*Transactions
	IncomingOrders  chan *Order //canais para conectar as rotinas simultaneas
	ProcessedOrders chan *Order
	Wait            *sync.WaitGroup
}

func NewBook(incomingOrders chan *Order, processedOrders chan *Order, waitGroup *sync.WaitGroup) *Book {
	return &Book{
		Order:           []*Order{},
		Transactions:    []*Transactions{},
		IncomingOrders:  incomingOrders,
		ProcessedOrders: processedOrders,
		Wait:            waitGroup,
	}

}

// Necessario trabalhar com fila
type OrderQueue []*Order

func (oq *OrderQueue) Add(order *Order) {
	*oq = append(*oq, order)

}

func (oq *OrderQueue) GetNextOrder() *Order {
	if len(*oq) == 0 {
		return nil

	}

	order := (*oq)[0]
	*oq = (*oq)[1:]
	return order

}

func (book *Book) Trade() {

	buyOrders := make(map[string]*OrderQueue) //Crio um asset em uma fila
	sellOrders := make(map[string]*OrderQueue)

	for order := range book.IncomingOrders {
		asset := order.Asset.ID

		if buyOrders[asset] == nil {
			buyOrders[asset] = &OrderQueue{}
		}

		if sellOrders[asset] == nil {
			sellOrders[asset] = &OrderQueue{}

		}

		if order.OrderType == "BUY" {
			book.TryMatch(order, sellOrders[asset], buyOrders[asset])

		} else {
			book.TryMatch(order, buyOrders[asset], sellOrders[asset])

		}

	}
}

func (book *Book) PricesMatch(order, matchOrder *Order) bool {
	if order.OrderType == "BUY" {
		return matchOrder.Price <= order.Price

	} else {
		return matchOrder.Price >= order.Price
	}

}

func (book *Book) CreateTransaction(incomingOrders, matchedOrder *Order) *Transactions {
	var buyOrder, sellOrder *Order

	if incomingOrders.OrderType == "BUY" {
		buyOrder, sellOrder = incomingOrders, matchedOrder

	} else {
		buyOrder, sellOrder = matchedOrder, incomingOrders
	}

	shares := incomingOrders.PendingShares

	if matchedOrder.PendingShares < shares {
		shares = matchedOrder.PendingShares

	}

  return NewTransactions(sellOrder, buyOrder, shares, matchedOrder.Price)

}

func (book *Book) ProcessTransactions(transaction *Transactions) {
  defer book.Wait.Done()
  transaction.Process()
  book.RecordTransaction(transaction)
  book.ProcessedOrders <- transaction.BuyOrder
  book.ProcessedOrders <- transaction.SellingOrder
  
}

func (book *Book) RecordTransaction(transaction *Transactions) {
  book.Transactions = append(book.Transactions, transaction)
  transaction.BuyOrder.Transactions = append(transaction.BuyOrder.Transactions, transaction)
  transaction.SellingOrder.Transactions = append(transaction.SellingOrder.Transactions, transaction)

}



func (book *Book) TryMatch(newOrder *Order, availableOrders, pendingOrders *OrderQueue) {
	for {
		possibleMatch := availableOrders.GetNextOrder()
		if possibleMatch == nil {
			break

		}

		if !book.PricesMatch(newOrder, possibleMatch) {
			availableOrders.Add(possibleMatch)
			break
		}

		if possibleMatch.PendingShares > 0 {
			matchedTransaction := book.CreateTransaction(newOrder, possibleMatch)
			book.ProcessTransactions(matchedTransaction)

			if possibleMatch.PendingShares > 0 {
				availableOrders.Add(possibleMatch)

			}

			if newOrder.PendingShares == 0 {
				break

			}

		}

	}

	if newOrder.PendingShares > 0 {
		pendingOrders.Add(newOrder)

	}

}
