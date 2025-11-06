package requests

import (
	client "broker/pkg/clients"
	"broker/pkg/order_history"
	"broker/pkg/structs"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrNoUser     = errors.New("user not found")
	ErrUserExists = errors.New("user already exists")
	ErrBadPass    = errors.New("invalid password")
	ErrDb         = errors.New("database error")
)

type DealForm struct {
	Ticker string `json:"ticker"`
	Price  int    `json:"price"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}

type RequestsRepo interface {
	Add(*DealForm, *client.Client) error
	DeleteByID(int) error
	GetAll() ([]*structs.Deal, error)
}

type RequestsRepoImpl struct {
	DB *sql.DB
	order_history order_history.OrderHistoryRepo
}

func NewRepository(db *sql.DB, order_history order_history.OrderHistoryRepo) *RequestsRepoImpl {
	return &RequestsRepoImpl{DB: db, order_history: order_history}
}

func (repo *RequestsRepoImpl) Add(deal *DealForm, clnt *client.Client) error {
	var isBuy int
	switch deal.Type {
	case "BUY":
		isBuy = 1
	case "SELL":
		isBuy = 0
	default:
		return errors.New(fmt.Sprintf("Incorrect type: %v", deal.Type))
	}

	_, err := repo.DB.Exec(
		"INSERT INTO requests (user_id, ticker, vol, price, is_buy) VALUES (?, ?, ?, ?, ?) ",
		clnt.ID, deal.Ticker, deal.Amount, deal.Price, isBuy,
	)
	if err != nil {
		return ErrDb
	}
	err = repo.order_history.Add(&structs.Deal{
		Ticker: deal.Ticker,
		Price: deal.Price,
		Type: deal.Type,
		Amount: deal.Amount,
		UserID: clnt.ID,
		IsBuy: isBuy,
	})
	if err != nil {
		return ErrDb
	}

	return nil
}

func (repo *RequestsRepoImpl) DeleteByID(id int) error {
	_, err := repo.DB.Exec("DELETE FROM requests WHERE id = ? ", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RequestsRepoImpl) GetAll() ([]*structs.Deal, error) {
	items := []*structs.Deal{}
	rows, err := repo.DB.Query("SELECT id, user_id, ticker, vol, price, is_buy FROM requests")
	if err != nil {
		return nil, ErrDb
	}
	defer rows.Close() // надо закрывать соединение, иначе будет течь
	for rows.Next() {
		dl := &structs.Deal{}
		err = rows.Scan(&dl.ID, &dl.UserID, &dl.Ticker, &dl.Amount, &dl.Price, &dl.IsBuy)
		if err != nil {
			return nil, ErrDb
		}
		items = append(items, dl)
	}
	return items, nil
}
