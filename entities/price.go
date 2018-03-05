package entities

//Price an Entity capturing details about price of a Commodity
// 	ID: TODO
// 	Currency: TODO
// 	Amount: TODO
// 	StartTime: TODO
// 	EndTimeType: TODO
// 	CommodityId: TODO
type Price struct {
	ID          string
	Currency    string
	Amount      int64
	StartTime   int64
	EndTime     int64
	Type        string
	CommodityID string
	Commodity   *Commodity
}
