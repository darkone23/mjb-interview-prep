package user

import (
	"context"
	"database/sql"
	"fmt"
	"mjb-interview-prep/internal/db"
	"mjb-interview-prep/models"

	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type service struct {
	conf db.DbConf
}

func NewService(conf db.DbConf) (*service, error) {
	return &service{conf: conf}, nil
}

type User struct {
	Name string `json:"username" xml:"username" binding:"required"`
	// server generated argon2 encoded password hash
	// why argon2 encoded? because we will use the encoded salt to compare
	Password string `json:"password" xml:"password" binding:"required"`
}

func (s *service) AddUser(u User) (string, error) {
	dbConn := s.conf.Sqlite.ConnectionUrl
	dbDriver := "sqlite3" // "postgres"

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
