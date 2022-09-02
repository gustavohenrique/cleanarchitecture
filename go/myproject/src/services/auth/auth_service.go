package auth

import (
	"context"

	"{{ .ProjectName }}/src/entities"
	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/repositories"
	"{{ .ProjectName }}/src/shared/conf"
	"{{ .ProjectName }}/src/shared/cryptus"
	"{{ .ProjectName }}/src/shared/jsonwebtoken"
	"{{ .ProjectName }}/src/shared/logger"
	"{{ .ProjectName }}/src/shared/uuid"
)

type AuthService struct {
	jwt    *jsonwebtoken.Jwt
}

func NewService(repositoryContainer repositories.RepositoryContainer) interfaces.IAuthService {
	config := conf.Get()
	return &AuthService{
		jwt: jsonwebtoken.New(jsonwebtoken.Config{
			Secret:     config.Auth.Jwt.Secret,
			Audience:   config.Auth.Jwt.Audience,
			Expiration: config.Auth.Jwt.Expiration,
		}),
	}
}

func (s *AuthService) FindClientSecretHash(ctx context.Context, clientID string) (string, error) {
	// should retrieve from storage/db
	found := "2a1c0aac2bdfec1a6a8fc6712ff75450b30b22308fc6aa6b418aec6a0dc66ea22c9a8053bd975a990bbbabc2c910d86604dc47f3f3006f638a868618ff54a899$a14dc4033796c40440b7c7a8f73c00a5111d8efc78e109c1c1c55c1e6ba5c53e"
	return found, nil
}

func (s *AuthService) GenerateOauth2Credentials(ctx context.Context, clientID string) (entities.Oauth2Entity, error) {
	id := clientID
	if len(id) != 36 {
		id = uuid.NewV4()
	}
	secret := cryptus.RandomSecretKey()
	hash := cryptus.HashIt(secret, cryptus.Salt())
	logger.Info("you should store this secret hash together clientID:", hash)
	result := entities.Oauth2Entity{
		ClientID:     id,
		ClientSecret: secret,
	}
	return result, nil
}

func (s *AuthService) CompareClientSecretAndHash(ctx context.Context, clientSecret, hash string) (bool, error) {
	return cryptus.ComparePasswordAndHash(clientSecret, hash)
}

func (s *AuthService) GenerateJwt(ctx context.Context, id string) (string, error) {
	return s.jwt.GenerateToken(id)
}

func (s *AuthService) ParseJwt(ctx context.Context, token string) (string, error) {
	return s.jwt.ParseToken(token)
}
