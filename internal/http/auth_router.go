package http

import (
	"context"
	"fmt"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/service/authservice"
	"github.com/labstack/echo/v4"
)

func (r *Server) authRouterGroup(group *echo.Group, authService *authservice.Auth) {
	authGroup := group.Group("/auth")

	authGroup.POST("/register", func(c echo.Context) error {
		fmt.Println("register")
		ctx := context.Background()

		var register api.RegisterRequest
		if err := BindAndValidate(c, &register); err != nil {
			return err
		}

		authService.RegisterUser(ctx, &register)

		return nil
	})

	authGroup.POST("/login", func(c echo.Context) error {
		authService.LoginUser()

		return nil
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
