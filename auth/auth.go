package auth

import (
	"fmt"
	"time"

	"github.com/radianhanggata/go-pkg/initializer"
)

type Token struct {
	AccessToken  TokenDetails
	RefreshToken TokenDetails
}

type TokenDetails struct {
	Token   string `json:"token"`
	UUID    string `json:"uuid"`
	Expires int64  `json:"expires"`
}

type AccessDetails struct {
	UUID   string
	UserID interface{}
}

func CreateToken(userid uint, ev *initializer.EV) (token *Token, err error) {
	token = &Token{}
	now := time.Now()

	// access token
	atExp := now.Add(time.Minute * 15).Unix()
	token.AccessToken, err = generateToken(grantTypeAccess, userid, atExp, ev.AccessSecret)
	if err != nil {
		return
	}

	// refresh token
	rtExp := now.Add(time.Hour * 24 * 7).Unix()
	token.RefreshToken, err = generateToken(grantTypeRefresh, userid, rtExp, ev.RefreshSecret)
	if err != nil {
		return
	}

	return
}

func TokenValid(ts, secretKey string) error {
	token, err := extractToken(ts, secretKey)
	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(ts, envkey string) (ad *AccessDetails, err error) {
	token, err := extractToken(ts, envkey)
	if err != nil {
		return nil, err
	}

	claims := getTokenMapClaims(token)

	if claims[keyGrantType].(string) != grantTypeAccess {
		return nil, fmt.Errorf("token is not access type")
	}

	ad = &AccessDetails{
		UUID:   claims[keyUUID].(string),
		UserID: claims[keyUserID],
	}

	return
}
