package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JWTConf struct {
	SigningKey  string
	ExpiresTime int64
}

type CustomClaims struct {
	ID      int64
	Name    string
	Creator string
	jwt.StandardClaims
}

type JWTClient struct {
	SigningKey  []byte
	ExpiresTime int64
}

func NewJWTClient(key string, expires int64) *JWTClient {
	return &JWTClient{
		SigningKey:  []byte(key),
		ExpiresTime: expires,
	}
}

func (j *JWTClient) NewClaims(userID int64, name string, creator string) *CustomClaims {
	return &CustomClaims{
		ID:      userID,
		Name:    name,
		Creator: creator,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "mab",
			ExpiresAt: time.Now().Unix() + j.ExpiresTime,
			NotBefore: time.Now().Unix() - 1000,
		},
	}
}

func (j *JWTClient) Create(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWTClient) Parse(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenExpired
}
