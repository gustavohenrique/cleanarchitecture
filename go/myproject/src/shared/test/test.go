package test

import (
	"{{ .ProjectName }}/mocks"
	"{{ .ProjectName }}/src/infra"
	"{{ .ProjectName }}/src/repositories"
	"{{ .ProjectName }}/src/services"
	"{{ .ProjectName }}/src/shared/conf"
	"testing"

	"github.com/golang/mock/gomock"
)

const DB ContextString = "db"

type ContextString string

type TestHelper struct {
	authServiceMock *mocks.MockAuthService
}

func NewTestHelper() *TestHelper {
	return &TestHelper{}
}

func (s *TestHelper) WithAuthService(mock *mocks.MockAuthService) *TestHelper {
	s.authServiceMock = mock
	return s
}

func (s *TestHelper) WithOauth2Mock(t *testing.T) *TestHelper {
	ctrl := gomock.NewController(t)
	authServiceMock := mocks.NewMockAuthService(ctrl)
	config := conf.Get().Auth.Jwt.Test
	clientID := config.ClientID
	/*
		clientSecret := config.ClientSecret
		clientSecretHash := config.ClientSecretHash
		authServiceMock.
			EXPECT().
			FindClientSecretHash(Any(), clientID).
			Return(clientSecretHash, nil)
		authServiceMock.
			EXPECT().
			CompareClientSecretAndHash(Any(), clientSecret, clientSecretHash).
			Return(true, nil)
	*/
	authServiceMock.
		EXPECT().
		ParseJwt(Any(), Any()).
		Return(clientID, nil)
	return s.WithAuthService(authServiceMock)
}

func (s *TestHelper) GetServiceContainer() services.ServiceContainer {
	serviceContainer := GetServiceContainer()
	if s.authServiceMock != nil {
		serviceContainer.AuthService = s.authServiceMock
	}
	return serviceContainer
}

func GetServiceContainer() services.ServiceContainer {
	infraContainer := infra.New()
	repositoryContainer := repositories.New(infraContainer)
	return services.New(repositoryContainer)
}

func Any() gomock.Matcher {
	return gomock.Any()
}
