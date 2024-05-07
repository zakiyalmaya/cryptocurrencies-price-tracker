package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/config"
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

	duration := config.RedisExpiration
	expirationTime := time.Now().Add(duration)
	claims := &model.AuthClaims{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to create token")
	}

	err = u.repos.Redcl.Set(context.Background(), config.PrefixKeyTokenRedis+user.Username, tokenString, duration).Err()
	if err != nil {
		log.Println("Failed to store token in Redis:", err.Error())
		return nil, fmt.Errorf("failed to store token in Redis")
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

func (u *userSvcImpl) Logout(username string) error {
	err := u.repos.Redcl.Del(context.Background(), config.PrefixKeyTokenRedis+username).Err()
	if err != nil {
		log.Println("Failed to delete token from Redis:", err.Error())
		return fmt.Errorf("failed to delete token from Redis")
	}

	return nil
}
