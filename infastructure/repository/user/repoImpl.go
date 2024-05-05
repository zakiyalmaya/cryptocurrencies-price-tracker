package user

import (
	"database/sql"

	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type userRepoImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepoImpl{
		db: db,
	}
}

func (u *userRepoImpl) Create(user *model.UserEntity) error {
	_, err := u.db.Exec("INSERT INTO users (name, email, username, password) VALUES (?, ?, ?, ?)",
		user.Name, user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepoImpl) Get(username string) (*model.UserEntity, error) {
	user := &model.UserEntity{}
	res := u.db.QueryRow("SELECT id, name, email, username, password FROM users WHERE username = ?", username)
	
	if err := res.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Password); err != nil {
		return nil, err
	}

	return user, nil
}
