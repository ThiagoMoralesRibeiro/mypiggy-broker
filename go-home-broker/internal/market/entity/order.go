package entity

type Order struct {
	Id            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int //Quantificarmos quantas ações ainda não foram vendidas ou compradas até a ordem fechar
	Price         float64
	OrderType     string
	Status        string
	Transactions  []*Transactions //A order é composta de transações
}

func NewOrder(id string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order {
	return &Order{
		Id:            id,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        "PENDING",
		Transactions:  []*Transactions{},
	}

}

func (o *Order) ApplyTrade(tradeShares int) {
	if tradeShares <= 0 || o.PendingShares <= 0 {
		return
	}

	if tradeShares > o.PendingShares {
		tradeShares = o.PendingShares
	}

	o.PendingShares -= tradeShares

	if o.PendingShares == 0 {
		o.Status = "DONE"
	}
}

func (o *Order) AddTransaction(transaction *Transactions) {
  o.Transactions = append(o.Transactions, transaction)
  
}
