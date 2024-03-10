package order_history

type OrderHistory struct {
	ID     int     `json:"-"`
	Time   int64   `json:"-"`
	UserID int     `json:"-"`
	Ticker string  `json:"ticker"`
	Vol    int     `json:"volume"`
	Price  float32 `json:"-"`
	Type   string  `json:"type"`
	Status string  `json:"status"`
}

// CREATE TABLE `orders_history` (
//     `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
//     `time` int NOT NULL,
//     `user_id` int,
//     `ticker` varchar(300) NOT NULL,
//     `vol` int NOT NULL,
//     `price` float not null,
//     `is_buy` boolean not null,
//     KEY user_id(user_id)
// );
