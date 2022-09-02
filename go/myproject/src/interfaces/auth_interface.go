package interfaces

import (
	"context"

	"{{ .ProjectName }}/src/entities"
)

type IAuthService interface {
	FindClientSecretHash(ctx context.Context, clientID string) (string, error)
	GenerateOauth2Credentials(ctx context.Context, clientID string) (entities.Oauth2Entity, error)
	CompareClientSecretAndHash(ctx context.Context, clientSecret, hash string) (bool, error)
	GenerateJwt(ctx context.Context, id string) (string, error)
	ParseJwt(ctx context.Context, token string) (string, error)
}
