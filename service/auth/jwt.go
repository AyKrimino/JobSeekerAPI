package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/AyKrimino/JobSeekerAPI/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userID int, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": strconv.Itoa(userID),
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	}

	return nil, nil, fmt.Errorf("invalid token")
}
