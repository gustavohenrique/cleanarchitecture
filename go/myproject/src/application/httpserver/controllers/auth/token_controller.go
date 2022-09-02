package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"{{ .ProjectName }}/src/entities"
	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/services"
	"{{ .ProjectName }}/src/shared/conf"
	log "{{ .ProjectName }}/src/shared/logger"
	"{{ .ProjectName }}/src/valueobjects"
)

type TokenController struct {
	config      *conf.Config
	authService interfaces.IAuthService
}

func NewTokenController(serviceContainer services.ServiceContainer) *TokenController {
	return &TokenController{
		authService: serviceContainer.GetAuthService(),
		config:      conf.Get(),
	}
}

func (h *TokenController) AddRoutesTo(group *echo.Group) {
	group.POST("", h.Create)
	group.GET("/credentials", h.GenerateCredentials)
}

func (h *TokenController) GenerateCredentials(c echo.Context) error {
	res := valueobjects.NewHttpResponse()
	// ctx := c.Request().Context()
	// credentials, _ := h.authService.GenerateOauth2Credentials(ctx, "")
	// fake data
	credentials := entities.Oauth2Entity{
		ClientID:     "f7f4048b-58d6-4c1b-9c6a-41e96831c17a",
		ClientSecret: "3hV5bwvcamCnqcCS239d2b83908d4373af829a0c27f34af0",
	}
	res.SetData(credentials)
	return c.JSON(http.StatusCreated, res)
}

func (h *TokenController) Create(c echo.Context) error {
	var req entities.Oauth2Entity
	res := valueobjects.NewHttpResponse()
	if err := c.Bind(&req); err != nil {
		res.SetError(err, "Cannot marshal JSON")
		return c.JSON(http.StatusBadRequest, res)
	}
	ctx := c.Request().Context()
	hash, err := h.authService.FindClientSecretHash(ctx, req.ClientID)
	if err != nil {
		res.SetError(err, "Account not found or disabled")
		return c.JSON(http.StatusUnauthorized, res)
	}
	match, err := h.authService.CompareClientSecretAndHash(ctx, req.ClientSecret, hash)
	if !match {
		log.Error(err)
		res.SetError(err, "wrong password")
		return c.JSON(http.StatusUnauthorized, res)
	}
	jwtToken, err := h.authService.GenerateJwt(ctx, req.ClientID)
	if err != nil {
		log.Error("Failed to create a token:", err)
		res.SetError(err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	res.SetData(jwtToken)
	return c.JSON(http.StatusOK, res)
}

func (h *TokenController) Verify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if h.config.Auth.Disabled {
			return next(c)
		}
		token := c.Request().Header.Get(h.config.Auth.Jwt.Header)
		if strings.Trim(token, " ") == "" {
			return c.JSON(http.StatusExpectationFailed, "Expected header " + h.config.Auth.Jwt.Header)
		}
		ctx := c.Request().Context()
		clientID, err := h.authService.ParseJwt(ctx, token)
		if err != nil {
			return c.JSON(http.StatusForbidden, err.Error())
		}
		c.Set(conf.CONTEXT_CLIENT_ID, clientID)
		return next(c)
	}
}
