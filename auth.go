package rest_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
)

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
