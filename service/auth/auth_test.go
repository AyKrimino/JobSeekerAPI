package auth_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/AyKrimino/JobSeekerAPI/service/auth"
	"github.com/golang-jwt/jwt/v5"
)

func TestHashPassword(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		password := "Pass1234"

		hashed, err := auth.HashPassword(password)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if hashed == "" {
			t.Error("expected a non-empty hashed string")
		}

		if hashed == password {
			t.Error("hashed password should not be the same as the original password")
		}
	})

	t.Run("Empty String", func(t *testing.T) {
		password := ""

		hashed, err := auth.HashPassword(password)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if hashed == "" {
			t.Error("expected a non-empty hashed string")
		}
	})
}

func TestComparePassword(t *testing.T) {
	password := "Pass1234"
	hashed, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	tests := []struct {
		name     string
		hashed   string
		plain    []byte
		expected bool
	}{
		{"Correct Password", hashed, []byte(password), true},
		{"Incorrect Password", hashed, []byte("wrongpass"), false},
		{"Empty Hashed Password", "", []byte(password), false},
		{"Empty Plain Password", hashed, []byte(""), false},
		{"Invalid Hash", "notahash", []byte(password), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := auth.ComparePassword(tc.hashed, tc.plain)
			if result != tc.expected {
				t.Errorf("expected %v, got %v for case: %s", tc.expected, result, tc.name)
			}
		})
	}
}

func TestCreateJWT(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		userID := 1
		secret := []byte("secret")

		token, err := auth.CreateJWT(userID, secret)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if token == "" {
			t.Error("expected a non-empty token string")
		}
	})

	t.Run("Empty Secret", func(t *testing.T) {
		userID := 1
		secret := []byte("")

		token, err := auth.CreateJWT(userID, secret)
		if err == nil {
			t.Fatal("expected an error")
		}

		if token != "" {
			t.Error("expected an empty token string")
		}
	})
}

func TestValidateJWT(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		userID := 1
		secret := []byte("secret")

		tokenString, err := auth.CreateJWT(userID, secret)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		token, claims, err := auth.ValidateJWT(tokenString, secret)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !token.Valid {
			t.Fatal("expected token to be valid")
		}

		claimUserID, ok := claims["userID"].(string)
		if !ok {
			t.Fatal("expected userID claim to be a string")
		}

		if claimUserID != strconv.Itoa(userID) {
			t.Fatalf("expected userID %d, got %s", userID, claimUserID)
		}

		_, ok = claims["exp"].(float64)
		if !ok {
			t.Fatal("expected exp claim to be a valid timestamp")
		}
	})

	t.Run("Invalid token string", func(t *testing.T) {
		secret := []byte("secret")
		tokenString := "Invalid token string"

		_, _, err := auth.ValidateJWT(tokenString, secret)
		if err == nil {
			t.Fatal("expected an error")
		}
	})

	t.Run("Empty token string", func(t *testing.T) {
		secret := []byte("secret")
		tokenString := ""

		_, _, err := auth.ValidateJWT(tokenString, secret)
		if err == nil {
			t.Fatal("expected an error")
		}
	})

	t.Run("Expired token", func(t *testing.T) {
		userID := 1
		secret := []byte("secret")

		token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"userID": strconv.Itoa(userID),
				"exp":    time.Now().Add(-time.Hour * 24).Unix(),
			},
		)

		tokenString, err := token.SignedString(secret)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		_, _, err = auth.ValidateJWT(tokenString, secret)
		if err == nil {
			t.Fatal("expected an error")
		}
	})
}
