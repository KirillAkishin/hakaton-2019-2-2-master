package stats

import (
	"broker/pkg/service"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrDb = errors.New("database error")
)

type StatsRepo interface {
	GetAllByTicker(string) ([]*Stat, error)
	Clean() error
	Add(*service.OHLCV) error
}

type StatsRepoImpl struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *StatsRepoImpl {
	return &StatsRepoImpl{DB: db}
}

func (repo *StatsRepoImpl) GetAllByTicker(ticker string) ([]*Stat, error) {
	err := repo.Clean()
	if err != nil {
		return nil, ErrDb
	}

	items := []*Stat{}
	rows, err := repo.DB.Query(
		"SELECT id, time, open, high, low, close, volume, ticker FROM stats WHERE ticker = ?", ticker)
	if err != nil {
		return nil, ErrDb
	}
	defer rows.Close() // надо закрывать соединение, иначе будет течь
	for rows.Next() {
		stt := &Stat{}
		err = rows.Scan(
			&stt.ID, &stt.Time, &stt.Open, &stt.High, &stt.Low, &stt.Close, &stt.Volume, &stt.Ticker)
		if err != nil {
			return nil, ErrDb
		}
		items = append(items, stt)
	}
	return items, nil
}

func (repo *StatsRepoImpl) Clean() error {
	t := time.Now().Unix()
	_, err := repo.DB.Exec("DELETE FROM stats WHERE time < ?", t-300)
	if err != nil {
		return err
	}
	return nil
}

func (repo *StatsRepoImpl) Add(ohlcv *service.OHLCV) error {
	_, err := repo.DB.Exec(
		"INSERT INTO stats (time, interval, open, high, low, close, volume, ticker) VALUES (?, ?, ?, ?, ?) ",
		ohlcv.Time, ohlcv.Interval, ohlcv.Open, ohlcv.High, ohlcv.Low, ohlcv.Close, ohlcv.Volume, ohlcv.Ticker,
	)
	if err != nil {
		return ErrDb
	}
	return nil
}
