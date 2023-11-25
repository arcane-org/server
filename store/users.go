package store

import (
	"server/models"

	"github.com/jmoiron/sqlx"
)

type Users struct {
	db *sqlx.DB
}

func NewUsersStore(db *sqlx.DB) *Users {
	return &Users{
		db: db,
	}
}

func (s *Users) Insert(u *models.User) error {
	_, err := s.db.NamedExec(
    "insert into users(id, email, username, avatar_url) values(:id, :email, :username, :avatar_url)",
		u,
	)
	return err
}

func (s *Users) GetById(id string) (u *models.User, err error) {
	u = new(models.User)
	err = s.db.Get(
		u,
		"select * from users where id=$1",
		id,
	)
	return
}

func (s *Users) GetByEmail(email string) (u *models.User, err error) {
	u = new(models.User)
	err = s.db.Get(
		u,
		"select * from users where email=$1",
		email,
	)
	return
}
