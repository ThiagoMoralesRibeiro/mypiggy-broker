package entity

import (
	"github.com/google/uuid"
	"time"
)

type Transactions struct {
	Id             string
	SellingOrder   *Order
	SellingOrderId string
	BuyOrder       *Order
	BuyOrderId     string
	Shares         int
	Price          float64
	Total          float64
	DateTime       time.Time
}

func NewTransactions(sellOrder *Order, BuyOrder *Order, shares int, price float64) *Transactions {
	return &Transactions{
		Id:             uuid.New().String(),
		SellingOrder:   sellOrder,
		SellingOrderId: uuid.New().String(),
		BuyOrder:       BuyOrder,
		BuyOrderId:     uuid.New().String(),
		Shares:         shares,
		Price:          price,
		Total:          0,
		DateTime:       time.Now(),
	}

}

func (transaction *Transactions) Process() {
	processor := NewOrderProcessor(transaction)
	processor.Process()

}
