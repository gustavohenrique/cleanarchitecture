package jsonwebtoken

import (
	"crypto/rsa"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Config struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
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
		jwtConfig.Issuer = "myproject"
	}
	return &Jwt{jwtConfig}
}

func (j *Jwt) GenerateToken(id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("Cannot generate token for empty string")
	}
	expiration := time.Second * time.Duration(3600)
	i, err := strconv.Atoi(j.jwtConfig.Expiration)
	if err == nil {
		expiration = time.Second * time.Duration(i)
	}
	claims := jwt.StandardClaims{
		Id:        id,
		Issuer:    j.jwtConfig.Issuer,
		Audience:  j.jwtConfig.Audience,
		Subject:   j.jwtConfig.Subject,
		ExpiresAt: time.Now().Add(expiration).Unix(),
		IssuedAt:  time.Now().Unix(),
	}
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = claims
	return t.SignedString(j.jwtConfig.PrivateKey)
}

func (j *Jwt) ParseToken(tokenString string) (string, error) {
	t, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.jwtConfig.PublicKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("Invalid token")
	}
	if claims.Valid() != nil {
		return "", fmt.Errorf("Expired token")
	}
	isCorrectContext := claims.VerifyAudience(j.jwtConfig.Audience, true)
	if !isCorrectContext {
		return "", fmt.Errorf("Invalid token")
	}
	return claims.Id, nil
}
