package controller

import (
	"os"
	"strconv"
	"time"

	"github.com/gildemberg-santos/process-event-go/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	ClientID string `json:"client_id"`
	Token    string `json:"token"`
	jwt.StandardClaims
}

type AuthRequest struct {
	ClientID string `json:"client_id"`
	SecretID string `json:"secret_id"`
}

func AuthMiddleware(c *gin.Context) {
	tknStr := c.GetHeader("Authorization")
	if tknStr == "" {
		c.JSON(401, gin.H{
			"message": "Authorization Required",
		})
		c.Abort()
		return
	}

	claims, err := validateToken(tknStr)
	if err != nil {
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("client_id", claims.ClientID)
	c.Next()
}

func Auth(c *gin.Context) {
	var authRequest AuthRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		c.Abort()
		return
	}

	credential := model.NewCredential()

	if authRequest.ClientID != credential.ClientID || authRequest.SecretID != credential.SecretID {
		c.JSON(401, gin.H{
			"message": "Invalid client_id or secret_id",
		})
		c.Abort()
		return
	}

	tknStr := authRequest.ClientID

	token, err := generateToken(tknStr)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func generateToken(tknStr string) (string, error) {
	minute, err := strconv.Atoi(os.Getenv("EXPIRATION_TIME"))
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(time.Duration(minute) * time.Minute)
	claims := &Claims{
		ClientID: tknStr,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(secretKey)
}

func validateToken(tknStr string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
