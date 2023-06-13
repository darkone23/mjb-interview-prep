package user

import (
	"context"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	// "log"
	"mjb-interview-prep/internal/db"
	"mjb-interview-prep/models"
)

type service struct {
	conf db.SqlConf
	pool db.DbConnectionPool
}

func NewService(conf db.SqlConf, maxConcurrency int) (service, error) {
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

func (s service) FindAllUsers() ([]UserDto, error) {
	db := s.pool.Acquire()
	defer s.pool.Release(db)

	users, err := models.Users().All(context.Background(), db)
	if err != nil {
		return []UserDto{}, err
	} else {
		var dtos = make([]UserDto, len(users))
		for i, u := range users {
			dtos[i] = FromModel(u)
		}
		return dtos, nil
	}
}

func (s service) FindUser(id int) (UserDto, error) {
	db := s.pool.Acquire()
	defer s.pool.Release(db)

	var dto UserDto
	var found *models.User

	found, err := models.FindUser(context.Background(), db, int64(id))
	if err == nil {
		dto = FromModel(found)
	}
	return dto, err
}

func (s service) AddUser(u UserDto) (string, error) {
	db := s.pool.Acquire()
	defer s.pool.Release(db)

	db_user := u.ToModel()
	var err = db_user.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		return "", fmt.Errorf("failed to insert: %s", err)
	}

	return fmt.Sprintf("%d", db_user.UserID), nil
}
