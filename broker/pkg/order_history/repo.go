package order_history

import (
	"broker/pkg/structs"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrDb = errors.New("database error")
)

type OrderHistoryRepo interface {
	GetAll() ([]*OrderHistory, error)
	Clean() error
	Add(*structs.Deal) error
}

type OrderHistoryRepoImpl struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *OrderHistoryRepoImpl {
	return &OrderHistoryRepoImpl{DB: db}
}

func (repo *OrderHistoryRepoImpl) Clean() error {
	t := time.Now().Unix()
	_, err := repo.DB.Exec("DELETE FROM orders_history WHERE time < ?", t-300)
	if err != nil {
		return err
	}
	return nil
}

func (repo *OrderHistoryRepoImpl) GetAll() ([]*OrderHistory, error) {
	err := repo.Clean()
	if err != nil {
		return nil, ErrDb
	}
	items := []*OrderHistory{}
	rows, err := repo.DB.Query("SELECT id, time, user_id, ticker, vol, price, is_buy FROM orders_history")
	if err != nil {
		return nil, ErrDb
	}
	defer rows.Close() // надо закрывать соединение, иначе будет течь
	for rows.Next() {
		hist := &OrderHistory{}
		var isBuy int
		err = rows.Scan(&hist.ID, &hist.Time, &hist.UserID, &hist.Ticker, &hist.Vol, &hist.Price, &isBuy)
		if err != nil {
			return nil, ErrDb
		}
		if isBuy == 1 {
			hist.Type = "BUY"
		} else {
			hist.Type = "SELL"
		}
		items = append(items, hist)
	}
	return items, nil
}

func (repo *OrderHistoryRepoImpl) Add(dl *structs.Deal) error {
	t := time.Now().Unix()
	_, err := repo.DB.Exec(
		"INSERT INTO orders_history (time, user_id, ticker, vol, price, is_buy) VALUES (?, ?, ?, ?, ?, ?) ",
		t, dl.ID, dl.Ticker, dl.Amount, dl.Price, dl.IsBuy,
	)
	if err != nil {
		return ErrDb
	}

	return nil
}

