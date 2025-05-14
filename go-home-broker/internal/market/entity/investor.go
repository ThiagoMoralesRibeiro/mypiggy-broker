package entity

type Investor struct {
	Id            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

type InvestorAssetPosition struct {
	AssetId string
	Shares  int
}

func NewInvestor(id, name string) *Investor {
	return &Investor{
		Id:            id,
		Name:          name,
		AssetPosition: []*InvestorAssetPosition{},
	}

}

func NewInvestorAssetPosition(id string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetId: id,
		Shares:  shares,
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)

}

func (i *Investor) AdjustAssetPosition(id string, qtdShares int) {
	assetPosition := i.GetAssetPosition(id)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(id, qtdShares))
	} else {
    assetPosition.AddShares(qtdShares)
  }
}

func (i *Investor) GetAssetPosition(id string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetId == id {
			return assetPosition

		}

	}
	return nil

}

func (iAsset *InvestorAssetPosition) AddShares(shares int){
  iAsset.Shares += shares;
}
