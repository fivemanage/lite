package http

import (
	"github.com/fivemanage/lite/internal/service/authservice"
	"github.com/labstack/echo/v4"
)

func (r *Server) authRouterGroup(group *echo.Group, authService *authservice.Auth) {
	authGroup := group.Group("/auth")

	authGroup.GET("/login", func(c echo.Context) error {
		url := authService.Login()

		return c.Redirect(302, url)
	})

	authGroup.GET("/callback/github", func(c echo.Context) error {
		code := c.QueryParam("code")

		token := authService.Callback(code)

		return c.JSON(200, echo.Map{
			"token":          token.AccessToken,
			"resfresh_token": token.RefreshToken,
			"expiry":         token.Expiry,
		})
	})
}
