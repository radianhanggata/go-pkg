package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

func generateToken(accessType string, userid uint, exp int64, envkey string) (td TokenDetails, err error) {
	td = TokenDetails{
		UUID:    uuid.NewV4().String(),
		Expires: exp,
	}

	atClaims := jwt.MapClaims{
		keyUUID:      td.UUID,
		keyUserID:    userid,
		keyExp:       td.Expires,
		keyGrantType: accessType,
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.Token, err = at.SignedString([]byte(envkey))
	if err != nil {
		return
	}

	return td, nil
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
	return token.Claims.(jwt.MapClaims)
}
