package middleware

import (

	"errors"
	"os"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Authorization field not found"})
			c.Abort()
			return
		}

		headerSlice := strings.Split(authHeader, " ")
		if len(headerSlice) != 2 || strings.ToLower(headerSlice[0]) != "bearer" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Invalid authorization parameters"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(headerSlice[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("tokenization method incompatible")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.IndentedJSON(http.StatusForbidden, gin.H{"error":"invalid token"})
			c.Abort()
			return 
		}
		
		c.Next()
		c.Abort()
	}
}

