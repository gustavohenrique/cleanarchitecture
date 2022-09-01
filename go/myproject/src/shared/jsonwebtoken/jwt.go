package jsonwebtoken

import (
	"crypto/rsa"
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Config struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	Secret     string
	Audience   string
	Subject    string
	Issuer     string
	Expiration string
}

type Jwt struct {
	jwtConfig Config
}

func New(jwtConfig Config) *Jwt {
	if jwtConfig.Audience == "" {
		jwtConfig.Audience = "web"
	}
	if jwtConfig.Issuer == "" {
		jwtConfig.Issuer = "webapp"
	}
	return &Jwt{jwtConfig}
}

func (j *Jwt) GenerateToken(id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("cannot generate token for empty string")
	}
	expiration := time.Second * time.Duration(3600)
	i, err := strconv.Atoi(j.jwtConfig.Expiration)
	if err == nil {
		expiration = time.Second * time.Duration(i)
	}
	claims := jwt.RegisteredClaims{
		ID:        id,
		Issuer:    j.jwtConfig.Issuer,
		Audience:  []string{j.jwtConfig.Audience},
		Subject:   j.jwtConfig.Subject,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = claims
	key := j.jwtConfig.Secret
	if strings.TrimSpace(key) != "" {
		return t.SignedString(key)
	}
	return t.SignedString(j.jwtConfig.PrivateKey)
}

func (j *Jwt) ParseToken(tokenString string) (string, error) {
	t, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		key := j.jwtConfig.Secret
		if strings.TrimSpace(key) != "" {
			return key, nil
		}
		return j.jwtConfig.PublicKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}
	if claims.Valid() != nil {
		return "", fmt.Errorf("expired token")
	}
	isCorrectContext := claims.VerifyAudience(j.jwtConfig.Audience, true)
	if !isCorrectContext {
		return "", fmt.Errorf("invalid token")
	}
	return claims.ID, nil
}
