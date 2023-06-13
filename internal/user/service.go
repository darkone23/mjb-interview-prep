package user

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type service struct {
	dbUser     string
	dbPassword string
}

func NewService(dbUser, dbPassword string) (*service, error) {
	if dbUser == "" {
		return nil, errors.New("dbUser was empty")
	}
	return &service{dbUser: dbUser, dbPassword: dbPassword}, nil
}

type User struct {
	Name     string
	Password string
}

func (s *service) AddUser(u User) (string, error) {
	dbConn := "data/db.sqlite3" // "postgres://admin:admin@localhost/test_repo?sslmode=disable"
	dbDriver := "sqlite3"       // "postgres"

	db, err := sql.Open(dbDriver, dbConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id string
	// TODO: SQL injection detected... need to use prepared statement
	// would be better to just import some ORM lib
	q := "INSERT INTO users (username, password) VALUES ('" + u.Name + "', '" + u.Password + "') RETURNING id"

	// cannot use query to insert...
	// cannot scan a nonexistent id...
	err = db.QueryRow(q).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}
