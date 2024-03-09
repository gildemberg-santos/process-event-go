package controller

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	Username   string `json:"username"`
	Token      string `json:"token"`
	Domain     string `json:"domain"`
	RemoteAddr string `json:"remote_addr"`
	UserAgent  string `json:"user_agent"`
	jwt.StandardClaims
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	tokenResponse := Claims{
		Token:      tknStr,
		Domain:     c.Request.Host,
		RemoteAddr: c.Request.RemoteAddr,
		UserAgent:  c.Request.UserAgent(),
	}

	claims, err := validateToken(tokenResponse)
	if err != nil {
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
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

	if authRequest.Username != os.Getenv("USER_NAME") || authRequest.Password != os.Getenv("USER_PASSWORD") {
		c.JSON(401, gin.H{
			"message": "Invalid username or password",
		})
		c.Abort()
		return
	}

	tokenResponse := Claims{
		Username:   authRequest.Username,
		Domain:     c.Request.Host,
		RemoteAddr: c.Request.RemoteAddr,
		UserAgent:  c.Request.UserAgent(),
	}

	token, err := generateToken(tokenResponse)
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

func generateToken(tokenResponse Claims) (string, error) {
	minute, err := strconv.Atoi(os.Getenv("EXPIRATION_TIME"))
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(time.Duration(minute) * time.Minute)
	claims := &Claims{
		Username:   tokenResponse.Username,
		Domain:     tokenResponse.Domain,
		RemoteAddr: tokenResponse.RemoteAddr,
		UserAgent:  tokenResponse.UserAgent,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(secretKey)
}

func validateToken(tokenResponse Claims) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenResponse.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	if claims.Domain != tokenResponse.Domain || claims.RemoteAddr != tokenResponse.RemoteAddr || claims.UserAgent != tokenResponse.UserAgent {
		return nil, jwt.ErrHashUnavailable
	}
	return claims, nil
}
