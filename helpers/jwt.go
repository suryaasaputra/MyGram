package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("secret")

func GenerateToken(id int, email, username string, age int) (string, error) {
	claims := jwt.MapClaims{
		"id":        id,
		"email":     email,
		"user_name": username,
		"age":       age,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString(secretKey)
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
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, errors.New("failed to claims or token not valid")
	}

	return token.Claims.(jwt.MapClaims), nil
}
