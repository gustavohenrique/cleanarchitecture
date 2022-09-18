package test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"{{ .ProjectName }}/mocks"
	"{{ .ProjectName }}/src/domain/gateways"
	"{{ .ProjectName }}/src/domain/services"
	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores"
	"{{ .ProjectName }}/src/interfaces"
)

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

func (s *TestHelper) GetServices() interfaces.IService {
	servicez := GetServices()
	if s.authServiceMock != nil {
		servicez.SetAuthService(s.authServiceMock)
	}
	return servicez
}

func GetServices() interfaces.IService {
	config := conf.Get()
	stores := datastores.With(config).New()
	gates := gateways.With(config).Inject(stores)
	return services.With(config).Inject(gates)
}

func Any() gomock.Matcher {
	return gomock.Any()
}
