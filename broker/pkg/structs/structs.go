package structs

type Deal struct {
	Ticker string `json:"ticker"`
	Price  int    `json:"price"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`
	ID     int
	UserID int
	IsBuy  int
}
