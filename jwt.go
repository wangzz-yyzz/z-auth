package z_auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// GenerateToken generate token string with the configuration
// config: Configuration created by NewConfiguration or DefaultConfiguration
// return: token string
func GenerateToken(config Configuration) (string, error) {
	// get the expiry time
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.ExpireTime) * time.Hour)

	// create the claims
	claims := Claims{
		config.UserName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Signer,
		},
	}

	// generate the token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.JwtSecret))

	return token, err
}

// ParseToken parse the token and return the claims. if the token is invalid, return nil
// token: token string
// config: Configuration created by NewConfiguration or DefaultConfiguration
// return: Claims
func ParseToken(token string, config Configuration) (*Claims, error) {
	// parse the token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})

	// if token is valid, return the claims
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	// if token is invalid, return nil
	return nil, err
}
