package http

import (
	"embed"
	nethttp "net/http"

	"github.com/fivemanage/lite/internal/service/authservice"
	"github.com/labstack/echo/v4"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
)

// go:embed ../../../web/dist
var webContent embed.FS

type Server struct {
	Engine *echo.Echo
}

func NewServer(authService *authservice.Auth) *Server {
	engine := echo.New()
	srv := &Server{
		Engine: engine,
	}

	srv.Engine.Validator = &CustomValidator{validator: validator.New()}

	srv.Engine.Use(middleware.Logger())
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
