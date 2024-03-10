package clients

import (
	"database/sql"
	"errors"
)

const BALANCE = 100500

var (
	ErrNoUser     = errors.New("user not found")
	ErrUserExists = errors.New("user already exists")
	ErrBadPass    = errors.New("invalid password")
	ErrDb         = errors.New("database error")
)

type ClientRepo interface {
	GetByID(int) (*Client, error)
	GetByName(string) (*Client, error)
	Register(*ClientForm) (*Client, error)
}

type ClientForm struct {
	Name string
	Pass string
}

type ClientRepoImpl struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *ClientRepoImpl {
	return &ClientRepoImpl{DB: db}
}

func (repo *ClientRepoImpl) GetByID(id int) (*Client, error) {
	post := &Client{}
	err := repo.DB.
		QueryRow("SELECT id, name, password, balance FROM clients WHERE id = ?", id).
		Scan(&post.ID, &post.Name, &post.Password, &post.Balance)
	if err == sql.ErrNoRows {
		return nil, ErrNoUser
	} else if err != nil {
		return nil, ErrDb
	}
	return post, nil
}

func (repo *ClientRepoImpl) GetByName(name string) (*Client, error) {
	post := &Client{}
	err := repo.DB.
		//insert into <название таблицы> ([<Имя столбца>, ... ]) values (<Значение>,...)
		QueryRow("SELECT id, name, password, balance FROM clients WHERE name = ?", name).
		Scan(&post.ID, &post.Name, &post.Password, &post.Balance)
	if err == sql.ErrNoRows {
		return nil, ErrNoUser
	} else if err != nil {
		return nil, ErrDb
	}
	return post, nil
}

func (repo *ClientRepoImpl) Register(form *ClientForm) (*Client, error) {
	u := &Client{}
	err := repo.DB.
		QueryRow("SELECT id, name, password, balance FROM clients WHERE name = ?", form.Name).
		Scan(&u.ID, &u.Name, &u.Password, &u.Balance)
	if err == nil {
		return nil, ErrUserExists
	} else if err != sql.ErrNoRows {
		return nil, ErrDb
	}

	result, err := repo.DB.Exec(
		"INSERT INTO clients (name, password, balance) VALUES (?, ?, ?) ",
		form.Name, form.Pass, BALANCE,
	)
	if err != nil {
		return nil, ErrDb
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, ErrDb
	}
	u = &Client{int(id), form.Name, form.Pass, BALANCE}
	return u, nil
}
