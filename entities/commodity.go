package entities

//Commodity an Entity capturing details about the actual product
// 	Id: TODO
// 	Name: TODO
// 	ParityId: TODO
// 	Quality: TODO
// 	Percentage: TODO
// 	Origin: TODO
// 	Quantity1: TODO
// 	Quantity2: TODO
// 	Unit: TODO
type Commodity struct {
	ID         string
	Name       string
	ParityID   string
	Parity     *Parity
	Quality    int64
	Percentage float32
	Origin     string
	Quantity1  int64
	Quantity2  int64
	Unit       string
}
