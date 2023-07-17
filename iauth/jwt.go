package iauth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

func generateToken(accessType string, value string, exp int64, envkey string, meta *TokenMetadata) (et string, err error) {
	uuid := uuid.NewV4().String()

	atClaims := jwt.MapClaims{
		keyUUID:      uuid,
		keyValue:     value,
		keyExp:       exp,
		keyGrantType: accessType,
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	et, err = at.SignedString([]byte(envkey))
	if err != nil {
		return
	}

	meta.UUID = uuid
	meta.Expires = exp
	meta.Value = value

	return et, nil
}

func extractToken(ts, envkey string) (*jwt.Token, error) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header[keyAlg])
		}
		return []byte(envkey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func getTokenMapClaims(token *jwt.Token) jwt.MapClaims {
	claims := token.Claims.(jwt.MapClaims)
	return claims
}

func extractError(err error) interface{} {
	jwterr := err.(*jwt.ValidationError)
	errdata := map[string]string{
		"code":    fmt.Sprint(jwterr.Errors),
		"message": jwterr.Error(),
	}
	return errdata
}
