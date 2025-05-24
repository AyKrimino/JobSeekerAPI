package auth

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userID int, userRole string, secret []byte) (string, error) {
	if len(secret) == 0 {
		return "", fmt.Errorf("secret key cannot be empty")
	}

	if !strings.EqualFold(userRole, "JobSeeker") && !strings.EqualFold(userRole, "Company") {
		return "", fmt.Errorf("invalid user role")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": strconv.Itoa(userID),
			"role":   userRole,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string, secret []byte) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	}

	return nil, nil, fmt.Errorf("invalid token")
}
