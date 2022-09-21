package interfaces

import (
	"context"

	"{{ .ProjectName }}/src/domain/entities"
)

type IAuthService interface {
	New(gateways IGateway) IAuthService
	FindClientSecretHash(ctx context.Context, clientID string) (string, error)
	GenerateOauth2Credentials(ctx context.Context, clientID string) (entities.Oauth2, error)
	CompareClientSecretAndHash(ctx context.Context, clientSecret, hash string) (bool, error)
	GenerateJwt(ctx context.Context, id string) (string, error)
	ParseJwt(ctx context.Context, token string) (string, error)
}
