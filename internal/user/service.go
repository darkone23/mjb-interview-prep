package user

import (
	"context"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"mjb-interview-prep/internal/db"
	"mjb-interview-prep/models"
)

type service struct {
	conf db.DbConf
	pool db.DbConnectionPool
}

func NewService(conf db.DbConf, maxConcurrency int) (service, error) {
	pool := db.NewDbConnectionPool(maxConcurrency)
	return service{
		conf: conf,
		pool: pool,
	}, nil
}

func (svc service) Open() {
	svc.pool.Open(svc.conf)
}

func (svc service) Close() {
	svc.pool.Close()
}

type User struct {
	Name string `json:"username" xml:"username" binding:"required"`
	// server generated argon2 encoded password hash
	// why argon2 encoded? because we will use the encoded salt to compare
	Password string `json:"password" xml:"password" binding:"required"`
}

func (s service) AddUser(u User) (string, error) {
	db := s.pool.Acquire()
	defer s.pool.Release(db)

	db_user := &models.User{
		Username: u.Name,
		Password: u.Password,
	}

	var err = db_user.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		return "", fmt.Errorf("failed to insert: %s", err)
	}

	return fmt.Sprintf("%d", db_user.UserID), nil
}
