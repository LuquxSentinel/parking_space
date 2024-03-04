package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UID   string
	Email string
	jwt.RegisteredClaims
}

var key string = ";lkfekr09ewir90ifojksafoiw09[pIEFR0P"

func NewJwt(uid, email string) (*string, error) {
	claims := &Claims{
		UID:   uid,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 2 * time.Hour).UTC()),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func ValidateToken(signedToken string) (uid *string, err error) {
	claims := new(Claims)

	token, err := jwt.ParseWithClaims(signedToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if !token.Valid {
		err = errors.New("invalid authorization header")
		return
	}

	uid = &claims.UID

	return
}
