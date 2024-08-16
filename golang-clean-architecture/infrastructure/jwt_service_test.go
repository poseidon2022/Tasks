package infrastructure

import (
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"errors"
	"github.com/stretchr/testify/suite"
	"golang-clean-architecture/domain"
)

type GenerateTokenTestSuite struct {
	suite.Suite
}

func (suite *GenerateTokenTestSuite) SetupSuite() {
	// Set the JWT secret for testing
	os.Setenv("JWT_SECRET", "testsecret")
}

func (suite *GenerateTokenTestSuite) TestGenerateTokenSuccess() {

	user := &domain.User{
		Email: "test@example.com",
		Role:  "admin",
	}
	token, err := GenerateToken(user)
	suite.NoError(err)
	suite.NotEmpty(token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("incompatible tokenization method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	suite.NoError(err)
	suite.True(parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	suite.True(ok)
	suite.Equal("test@example.com", claims["email"])
	suite.Equal("admin", claims["role"])

	exp, ok := claims["exp"].(float64)
	suite.True(ok)
	suite.True(int64(exp) > time.Now().Unix())
}

func TestGenerateTokenTestSuite(t *testing.T) {
	suite.Run(t, new(GenerateTokenTestSuite))
}
