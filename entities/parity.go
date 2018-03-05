package entities

//Parity an entity associated with a location [AC] to provide more context
// 	ID: TODO
// 	Incoterm: TODO
// 	LocationID: TODO
type Parity struct {
	ID         string
	Incoterm   string
	LocationID string
	Location   *Location
}
