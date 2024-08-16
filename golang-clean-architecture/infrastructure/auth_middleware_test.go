package infrastructure

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
    "time"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/suite"
    "golang-clean-architecture/domain"
)

type AuthMiddlewareTestSuite struct {
    suite.Suite
    router *gin.Engine
}

func (suite *AuthMiddlewareTestSuite) SetupSuite() {
    // Set up Gin router and middleware
    gin.SetMode(gin.TestMode)
    suite.router = gin.New()
    suite.router.Use(AuthMiddleWare())
    // Set up a dummy handler for testing
    suite.router.GET("/protected", func(c *gin.Context) {
        user, exists := c.Get("AuthorizedUser")
        if exists {
            c.JSON(http.StatusOK, user)
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
        }
    })
}

func (suite *AuthMiddlewareTestSuite) TestMissingAuthorizationHeader() {
    req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
    resp := httptest.NewRecorder()

    suite.router.ServeHTTP(resp, req)

    suite.Equal(http.StatusBadRequest, resp.Code)
    suite.Contains(resp.Body.String(), "authorization header not found")
}

func (suite *AuthMiddlewareTestSuite) TestInvalidAuthorizationHeaderFormat() {
    req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
    req.Header.Set("Authorization", "InvalidToken")
    resp := httptest.NewRecorder()

    suite.router.ServeHTTP(resp, req)

    suite.Equal(http.StatusBadRequest, resp.Code)
    suite.Contains(resp.Body.String(), "bearer token not found")
}

func (suite *AuthMiddlewareTestSuite) TestInvalidJWTToken() {
    req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
    req.Header.Set("Authorization", "Bearer invalidtoken")
    resp := httptest.NewRecorder()

    suite.router.ServeHTTP(resp, req)

    suite.Equal(http.StatusBadRequest, resp.Code)
    suite.Contains(resp.Body.String(), "invalid token")
}

func (suite *AuthMiddlewareTestSuite) TestValidJWTToken() {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": "test@example.com",
        "role":  "admin",
        "exp":   time.Now().Add(time.Hour * 1).Unix(),
    })
    tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

    req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
    req.Header.Set("Authorization", "Bearer "+tokenString)
    resp := httptest.NewRecorder()

    suite.router.ServeHTTP(resp, req)

    suite.Equal(http.StatusOK, resp.Code)

    var user domain.AuthenticatedUser
    err := json.Unmarshal(resp.Body.Bytes(), &user)
    suite.NoError(err)
    suite.Equal("test@example.com", user.Email)
    suite.Equal("admin", user.Role)
}

func TestAuthMiddlewareTestSuite(t *testing.T) {
    suite.Run(t, new(AuthMiddlewareTestSuite))
}