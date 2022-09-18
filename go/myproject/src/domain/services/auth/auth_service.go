package auth

import (
	"context"

	"{{ .ProjectName }}/src/domain/entities"
	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/shared/jsonwebtoken"
)

type AuthService struct {
	jwt    *jsonwebtoken.Jwt
	config *conf.Config
}

func With(config *conf.Config) interfaces.IAuthService {
	return &AuthService{
		config: config,
	}
}

func (s *AuthService) New(gateways interfaces.IGateway) interfaces.IAuthService {
	s.jwt = jsonwebtoken.New(jsonwebtoken.Config{
		Secret:     s.config.Auth.Jwt.Secret,
		Audience:   s.config.Auth.Jwt.Audience,
		Expiration: s.config.Auth.Jwt.Expiration,
	})
	return s
}

func (s *AuthService) FindClientSecretHash(ctx context.Context, clientID string) (string, error) {
	// should retrieve from storage/db
	found := "2a1c0aac2bdfec1a6a8fc6712ff75450b30b22308fc6aa6b418aec6a0dc66ea22c9a8053bd975a990bbbabc2c910d86604dc47f3f3006f638a868618ff54a899$a14dc4033796c40440b7c7a8f73c00a5111d8efc78e109c1c1c55c1e6ba5c53e"
	return found, nil
}

func (s *AuthService) GenerateOauth2Credentials(ctx context.Context, clientID string) (entities.Oauth2, error) {
	oauth2 := entities.NewOauth2(clientID)
	return oauth2, nil
}

func (s *AuthService) CompareClientSecretAndHash(ctx context.Context, clientSecret, hash string) (bool, error) {
	oauth2 := entities.Oauth2{
		ClientSecret: clientSecret,
		Hash:         hash,
	}
	return oauth2.ComparePasswordAndHash()
}

func (s *AuthService) GenerateJwt(ctx context.Context, id string) (string, error) {
	return s.jwt.GenerateToken(id)
}

func (s *AuthService) ParseJwt(ctx context.Context, token string) (string, error) {
	return s.jwt.ParseToken(token)
}
