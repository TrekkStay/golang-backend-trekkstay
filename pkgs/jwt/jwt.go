package jwt

import (
	"encoding/json"
	"fmt"
	"time"
	"trekkstay/pkgs/jwt/constant"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{secret: secret}
}

func (j *JWT) Generate(payload map[string]interface{}, expiry int) (map[string]interface{}, error) {
	jwtPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, constant.ErrCannotMarshalPayload
	}

	data := constant.JWTPayload{}
	if err := json.Unmarshal(jwtPayload, &data); err != nil {
		return nil, err
	}

	// Generate the JWT token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": data.UserId,                                                // Set the user id in the token
		"exp":     time.Now().Add(time.Second * time.Duration(expiry)).Unix(), // Set the expiry time
		"iat":     time.Now().Unix(),                                          // Set the token creation time
	})

	// Sign the token with the secret key
	tokenString, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, constant.ErrEncodingToken
	}

	// Create a token object with the generated token, expiry, and creation time
	token := map[string]interface{}{
		"token":      tokenString,
		"expiry":     expiry,
		"created_at": time.Now(),
	}

	return token, nil
}

func (j *JWT) Validate(tokenString string) (*constant.JWTPayload, error) {
	fmt.Println("TOKENNNNNNNNN", tokenString)
	fmt.Println("SECRET", j.secret)
	// Parse the token with the secret key.
	secretKey := []byte(j.secret)
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, constant.ErrCannotUnmarshalToken
	}

	// Check if the token is valid.
	if jwtToken.Valid {
		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			return nil, constant.ErrInvalidToken
		}

		return &constant.JWTPayload{
			UserId: claims["user_id"].(string),
		}, nil
	}

	return nil, constant.ErrInvalidToken
}
