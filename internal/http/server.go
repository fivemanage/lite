package http

import (
	"embed"
	nethttp "net/http"

	"github.com/fivemanage/lite/internal/service/authservice"
	"github.com/labstack/echo/v4"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
)

// This embed shit shouldn't be this far away.
// When we are ready, we can either build the React code directly into the server or root folder
// or just copy it in the Dockerfile.

// go:embed ../../../web/dist
var webContent embed.FS

type Server struct {
	Engine *echo.Echo
}

// TODO: Add sentry for monitoring. There should be an opt-out option.
func NewServer(authService *authservice.Auth) *Server {
	engine := echo.New()
	srv := &Server{
		Engine: engine,
	}

	srv.Engine.Validator = &CustomValidator{validator: validator.New()}

	// TODO: lets use 'logrus' for logging

	// srv.Engine.Use(middleware.Logger())
	srv.Engine.Use(middleware.Recover())

	srv.Engine.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "web/dist",
		Filesystem: nethttp.FS(webContent),
		HTML5:      true,
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/api/auth/login"
		},
	}))

	apiGroup := srv.Engine.Group("/api")

	srv.authRouterGroup(apiGroup, authService)
	srv.imageRouterGroup(apiGroup)

	return srv
}
