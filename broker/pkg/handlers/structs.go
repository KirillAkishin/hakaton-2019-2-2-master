package handlers

type DealForm struct {
	Ticker string `json:"ticker"`
	Price  int    `json:"price"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}

type IdForm struct {
	Id int `json:"id"`
}
