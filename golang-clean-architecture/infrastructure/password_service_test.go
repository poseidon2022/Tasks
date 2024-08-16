package infrastructure

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"golang-clean-architecture/domain"
	"golang.org/x/crypto/bcrypt"
)

type PasswordTestSuite struct {
	suite.Suite
}

func (suite *PasswordTestSuite) TestHashPasswordSuccess() {
	password := "securepassword"
	hashedPassword, err := HashPassword(password)

	suite.NoError(err)
	suite.NotEmpty(hashedPassword)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	suite.NoError(err)
}

func (suite *PasswordTestSuite) TestHashPasswordFailure() {

	password := "securepassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost+1)
	suite.Error(err)
	suite.Empty(hashedPassword)
}



func (suite *PasswordTestSuite) TestComparePasswordsSuccess() {
	password := "securepassword"

	hashedPassword, err := HashPassword(password)
	suite.NoError(err)
	existingUser := &domain.User{Password: string(hashedPassword)}
	userInfo := &domain.User{Password: password}
	err = ComparePasswords(existingUser, userInfo)

	// Assertions
	suite.NoError(err)
}

func (suite *PasswordTestSuite) TestComparePasswordsFailure() {
	// Mock existing user with a hashed password
	existingUser := &domain.User{Password: "$2a$10$7eqZslQMGfMLyY/Si7D6deBuoJ7z/ncm8lf9WzlB9.Pz5IH.mBz3O"} // hashed "securepassword"
	userInfo := &domain.User{Password: "wrongpassword"}

	err := ComparePasswords(existingUser, userInfo)

	// Assertions
	suite.Error(err)
	suite.EqualError(err, "passwords don't match")
}

func TestPasswordTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}
