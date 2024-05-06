package user

import (
	"database/sql"
	"fmt"
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
			return nil, fmt.Errorf("invalid username")
		}
		return nil, err
	}

	if user.Password != auth.Password {
		return nil, fmt.Errorf("invalid password")
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
		return nil, fmt.Errorf("failed to create token")
	}

	return &model.AuthResponse{
		Name:     user.Name,
		Username: user.Username,
		Token:    tokenString,
	}, nil
}

func (u *userSvcImpl) Register(user *model.UserEntity) error {
	if err := u.repos.User.Create(user); err != nil {
		return fmt.Errorf("error when register user to database")
	}

	return nil
}
