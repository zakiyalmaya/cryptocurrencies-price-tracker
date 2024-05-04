package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserEntity struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
