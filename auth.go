package rest_api

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
)

// ApiKeyAuth метод для аутентифкации Basic
func ApiKeyAuth(Verify func(apiKey string) bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("api-key")
		if apiKey == "" {
			ErrorHandler(NewUnauthorizedError(errors.New("authorization api-key not found"), "Authorization api-key not found"), c)
			return
		}

		if !Verify(apiKey) {
			ErrorHandler(NewUnauthorizedError(errors.New("authorization api-key is invalid"), "Authorization api-key is invalid"), c)
			return
		}

		c.Next()
	}
}

// BasicAuth метод для аутентифкации Basic
func BasicAuth(Verify func(userName string, password string) bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			ErrorHandler(NewUnauthorizedError(errors.New("authorization token not found"), "Authorization token not found"), c)
			return
		}
		tokenHeader := strings.Replace(authHeader, "Basic ", "", -1)
		decodeToken, err := base64.StdEncoding.DecodeString(tokenHeader)
		if err != nil {
			ErrorHandler(NewUnauthorizedError(errors.New("authorization token is invalid"), "Authorization token not found"), c)
			return
		}
		claims := strings.Split(string(decodeToken), `:`)
		if len(claims) > 0 && !Verify(claims[0], claims[1]) {
			ErrorHandler(NewUnauthorizedError(errors.New("authorization token is invalid"), "Authorization token is invalid"), c)
			return
		}
		marshal, _ := json.Marshal(&struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Username: claims[0],
			Password: claims[1],
		})
		c.Set("user", marshal)
		c.Next()
	}
}

// JWTAuth Метод для аутентификации по JWT
func JWTAuth(GetClaim func() jwt.Claims) func(c *gin.Context) {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			ErrorHandler(NewUnauthorizedError(errors.New("authorization token not found"), "Authorization token not found"), c)
			return
		}
		tokenHeader := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := GetClaim()
		tkn, err := jwt.ParseWithClaims(tokenHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			fmt.Println(err.Error())
			ErrorHandler(NewUnauthorizedError(err, "Authorization token Invalid"), c)
			return
		}

		if !tkn.Valid {
			ErrorHandler(NewUnauthorizedError(errors.New("authorization token Invalid"), "Authorization token Invalid"), c)
			return
		}

		marshal, err := json.Marshal(claims)
		if err != nil {
			ErrorHandler(NewUnauthorizedError(errors.New("authorization token Invalid"), "Authorization token Invalid"), c)
			return
		}
		c.Set("user", marshal)
		c.Next()
	}
}
