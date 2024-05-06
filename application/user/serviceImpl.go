package user

import (
	"database/sql"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type userSvcImpl struct {
	repos *repository.Repositories
}

func NewUserService(repos *repository.Repositories) UserService {
	return &userSvcImpl{
		repos: repos,
	}
}

func (u *userSvcImpl) Login(auth *model.AuthRequest) (*model.AuthResponse, error) {
	user, err := u.repos.User.Get(auth.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid username")
		}
		return nil, err
	}

	if user.Password != auth.Password {
		return nil, errors.New("invalid password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &model.AuthClaims{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("cryptocurrencies-price-tracker-secret"))
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Name:     user.Name,
		Username: user.Username,
		Token:    tokenString,
	}, nil
}

func (u *userSvcImpl) Register(user *model.UserEntity) error {
	if err := u.repos.User.Create(user); err != nil {
		return err
	}

	return nil
}
