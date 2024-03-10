package stats

type Stat struct {
	ID       int     `json:"-"`
	Time     int64   `json:"time"`
	Interval int     `json:"-"`
	Open     float32 `json:"open"`
	High     float32 `json:"high"`
	Low      float32 `json:"low"`
	Close    float32 `json:"close"`
	Volume   int     `json:"volume"`
	Ticker   string  `json:"-"`
}

// `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
// `time` int,
// `interval` int,
// `open` float,
// `high` float,
// `low` float,
// `close` float,
// `volume` int,
// `ticker` varchar(300),
// KEY id(id)
// );
