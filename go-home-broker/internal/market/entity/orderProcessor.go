package entity

type OrderProcessor struct {
	Transactions *Transactions
}

func NewOrderProcessor(transaction *Transactions) *OrderProcessor {
	return &OrderProcessor{
		Transactions: transaction,
	}

}

//Etapas para processar uma Ordem
// - Calcular a quantidade de cotas
// - Mudar a posição de quem compra e quem vende. (Ex: Se uma pessoa que possui 100 e vende 50, agora ele possui somente 50 cotas)
// - Alterar o processamento da Ordem, alterando shares, status etc

func (orderProcessor *OrderProcessor) Process() {
	shares := orderProcessor.CalculateShares()
	orderProcessor.UpdatePosition(shares)
	orderProcessor.UpdateOrders(shares)
	orderProcessor.Transactions.Total = float64(shares) * orderProcessor.Transactions.Price

}

func (orderProcessor *OrderProcessor) CalculateShares() int {
	availableShares := orderProcessor.Transactions.Shares

	if orderProcessor.Transactions.BuyOrder.PendingShares < availableShares {
		availableShares = orderProcessor.Transactions.BuyOrder.PendingShares

	}

	if orderProcessor.Transactions.SellingOrder.PendingShares < availableShares {
		availableShares = orderProcessor.Transactions.SellingOrder.PendingShares
	}

	return availableShares

}

func (orderProcessor *OrderProcessor) UpdatePosition(shares int) {
	//Subtraindo uma cota da carteira do usuario
	orderProcessor.Transactions.SellingOrder.Investor.AdjustAssetPosition(orderProcessor.Transactions.SellingOrder.Asset.ID, -shares)
	//Adicionando uma cota a carteira do usuario
	orderProcessor.Transactions.BuyOrder.Investor.AdjustAssetPosition(orderProcessor.Transactions.BuyOrder.Asset.ID, shares)

}

func (orderProcessor *OrderProcessor) UpdateOrders(shares int) {
	orderProcessor.Transactions.SellingOrder.ApplyTrade(shares)
	orderProcessor.Transactions.BuyOrder.ApplyTrade(shares)

}
