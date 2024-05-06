package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

func AuthMiddleware(redcl *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.HTTPErrorResponse("Missing Authorization header"))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("cryptocurrencies-price-tracker-secret"), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.HTTPErrorResponse("Invalid or expired token"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.HTTPErrorResponse("Invalid token claims"))
			return
		}

		username := claims["username"].(string)
		c.Set("username", username)
		userID := claims["userId"].(float64)
		c.Set("userId", userID)

		tokenCache, err := redcl.Get(context.Background(), "jwt-token-"+username).Result()
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.HTTPErrorResponse("Invalid or expired token"))
			return
		}
		
		if tokenCache != tokenString {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.HTTPErrorResponse("Invalid or expired token"))
		}

		c.Next()
	}
}
