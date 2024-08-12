package middleware

import (

	"errors"
	models "task06/models"
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

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.IndentedJSON(http.StatusForbidden, gin.H{"error":"error while extracting claims"})
			c.Abort()
			return
		}

		user_id, ok := claims["user_id"].(string)

		if !ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user_id not found in the token"})
			c.Abort()
			return 
		}

		email, ok := claims["email"].(string)

		if !ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "email not found in the token"})
			c.Abort()
			return 
		}
		

		role, ok := claims["role"].(string)
		if !ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "role not found in the token"})
			c.Abort()
			return 
		}
		
		c.Set("AuthorizedUser", &models.AuthenticatedUser{
			ID : user_id,
			Role : role,
			Email : email,
		})

		c.Next()
	}
}

