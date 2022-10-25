package jwt

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"
)

const (
	issuerKey = "iss"
	expKey    = "exp"
	subKey    = "sub"

	issuer = "pad"
	key    = "9e28d10293fd9f37299e0849a9720fca"
)

func CreateToken(userID uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		subKey:    strconv.FormatUint(userID, 10),
		issuerKey: issuer,
	})

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ExtractUserFromToken(token string) (uint64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}

		return key, nil
	})
	if err != nil {
		return 0, err
	}

	var sub uint64
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if claims[issuerKey] != issuer {
			return 0, errors.New("issuers mismatch")
		}

		if sub, ok = claims[subKey].(uint64); !ok {
			return 0, errors.New("missing sub claim")
		}
	}

	return sub, nil
}
