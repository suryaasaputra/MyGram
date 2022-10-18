package helpers

import (
	"errors"
	"mygram/config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(id int, email, username string, age int) (string, error) {
	claims := jwt.MapClaims{
		"id":        id,
		"email":     email,
		"user_name": username,
		"age":       age,
		"exp":       time.Now().Add(time.Minute * 10).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	headerToken := ctx.Request.Header.Get("Authorization")
	haveBearer := strings.HasPrefix(headerToken, "Bearer")
	if !haveBearer {
		return nil, errors.New("no bearer in request header for authorization")
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("failed to parse token, expected jwt.SigningMethodHMAC")
		}
		return []byte(config.SECRET_KEY), nil
	})

	if err != nil {
		return nil, errors.New("failed to parse token, token invalid or expired")
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, errors.New("failed to claims or token not valid")
	}

	return token.Claims.(jwt.MapClaims), nil
}
