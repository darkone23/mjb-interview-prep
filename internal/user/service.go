package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mjb-interview-prep/models"

	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	Name string `json:"username" xml:"username" binding:"required"`
	// server generated argon2 encoded password hash
	// why argon2 encoded? because we will use the encoded salt to compare
	Password string `json:"password" xml:"password" binding:"required"`
}

func (s *service) AddUser(u User) (string, error) {
	dbConn := "data/db.sqlite3" // "postgres://admin:admin@localhost/test_repo?sslmode=disable"
	dbDriver := "sqlite3"       // "postgres"

	db, err := sql.Open(dbDriver, dbConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db_user := &models.User{
		Username: u.Name,
		Password: u.Password,
	}

	err = db_user.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		return "", fmt.Errorf("failed to insert: %w", err)
	}

	return fmt.Sprintf("%d", db_user.UserID), nil
}
