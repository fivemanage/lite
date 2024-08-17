package http

import (
	"embed"
	nethttp "net/http"

	"github.com/fivemanage/lite/internal/service/authservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// go:embed ../../../web/dist
var webContent embed.FS

type Server struct {
	Engine *echo.Echo
}

type GithubEmail struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility"` // "public" or "private"
}

func NewServer(authService *authservice.Auth) *Server {
	server := echo.New()
	// server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	srv := &Server{
		Engine: server,
	}

	server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "web/dist",
		Filesystem: nethttp.FS(webContent),
		HTML5:      true,
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/api/auth/login"
		},
	}))

	apiGroup := server.Group("/api")
	apiGroup.GET("/auth/login", func(c echo.Context) error {
		url := authService.Login()

		return c.Redirect(302, url)
	})

	apiGroup.GET("/auth/callback/github", func(c echo.Context) error {
		code := c.QueryParam("code")

		accessToken := authService.Callback(code)

		return c.JSON(200, echo.Map{
			"token": accessToken,
		})
	})

	return srv
}
