package iauth

import (
	"time"

	"github.com/radianhanggata/go-pkg/ictx"
	ictx1 "github.com/radianhanggata/go-pkg/ictx"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type TokenMetadata struct {
	UUID    string
	Value   string
	Expires int64
}

func CreateToken(value string, ev *ictx1.EV) (pair *TokenPair, atmd, rtmd *TokenMetadata, err error) {
	pair = &TokenPair{}
	now := time.Now()

	// access token
	atExp := now.Add(time.Minute * 15).Unix()
	atmd = &TokenMetadata{}
	pair.AccessToken, err = generateToken(grantTypeAccess, value, atExp, ev.AccessSecret, atmd)
	if err != nil {
		return nil, nil, nil, ictx.ErrorUnauthorized.Embed(extractError(err))
	}

	// refresh token
	rtExp := now.Add(time.Hour * 24 * 7).Unix()
	rtmd = &TokenMetadata{}
	pair.RefreshToken, err = generateToken(grantTypeRefresh, value, rtExp, ev.RefreshSecret, rtmd)
	if err != nil {
		return nil, nil, nil, ictx.ErrorUnauthorized.Embed(extractError(err))
	}

	return
}

func TokenValid(ts, secretKey string) error {
	token, err := extractToken(ts, secretKey)
	if err != nil {
		return ictx.ErrorUnauthorized.Embed(extractError(err))
	}

	if !token.Valid {
		return ictx.ErrorUnauthorized.Embed(extractError(err))
	}
	return nil
}

func ExtractTokenMetadata(ts, envkey string) (ad *TokenMetadata, err error) {
	tokenJwt, err := extractToken(ts, envkey)
	if err != nil {
		return nil, ictx.ErrorUnauthorized.Embed(extractError(err))
	}

	claims := getTokenMapClaims(tokenJwt)

	ad = &TokenMetadata{
		UUID:    claims[keyUUID].(string),
		Value:   claims[keyValue].(string),
		Expires: int64((claims[keyExp].(float64))),
	}

	return
}
