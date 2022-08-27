package claim

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	jwt.StandardClaims
	ID int `json:"id"`
}

// Generate a new token using the standard functions in the package https://github.com/dgrijalva/jwt-go
func (c *Claim) GetToken(signingString string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(signingString))
}

// Receives a token, validate, and retrieve claims
func GetFromToken(tokenString, signingString string) (*Claim, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(signingString), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	iID, ok := claim["id"]
	if !ok {
		return nil, errors.New("invalid user id - not found")
	}

	id, ok := iID.(float64)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	return &Claim{ID: int(id)}, nil
}
