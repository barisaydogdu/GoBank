package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Stroge interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

// Aşağıdaki fonksiyon bu struct'ı döner
type PostgressStore struct {
	db *sql.DB
}

func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}
func (s *PostgressStore) createAccountTable() error {
	query := `create table  if not exists account(
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	//Veritabanına bir sinyal gönderir ve cevap alır
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) CreateAccount(*Account) error {
	return nil
}

func (s *PostgressStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgressStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgressStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
