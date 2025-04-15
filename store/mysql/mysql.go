package mysql

import (
	"database/sql"
	"fmt"
	"third-party-int/store"
)

func NewDB() (*sql.DB, error) {
	return sql.Open("mysql", "myuser:mypassword@tcp(localhost:3306)/mydb")
}

type MysqlStore struct {
	db *sql.DB
}

func NewJokeStore(db *sql.DB) *MysqlStore {
	return &MysqlStore{db: db}
}

func (s *MysqlStore) StoreJoke(joke *store.Joke) error {
	if s.db == nil {
		return fmt.Errorf("DB connection is nil")
	}
	_, err := s.db.Exec("INSERT INTO jokes (type, setup, punchline) VALUES (?, ?, ?)", joke.Types, joke.Setup, joke.Punchline)
	return err
}
